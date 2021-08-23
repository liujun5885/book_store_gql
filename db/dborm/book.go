package dborm

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/graph/model"
)

type Book struct {
	DB *pg.DB
}

func (o *Book) SearchBooks(keyword string, pageCursor model.PageCursor) (*model.SearchBooksResponse, error) {
	var books []*model.Book
	count, err := o.DB.Model(&books).Where("title ILIKE ?", fmt.Sprintf("%%%s%%", keyword)).Offset(
		(pageCursor.Page - 1) * pageCursor.PageSize).Limit(pageCursor.PageSize).SelectAndCount()
	if err != nil {
		return nil, err
	}

	totalPage := count / pageCursor.PageSize
	if count%pageCursor.PageSize > 0 {
		totalPage += 1
	}

	return &model.SearchBooksResponse{
		PageInfo: &model.PageInfo{
			TotalItems: count,
			TotalPages: totalPage,
			Page:       pageCursor.Page,
			PageSize:   pageCursor.PageSize,
		},
		Books: model.ReshapeBooks(books),
	}, err
}

func (o *Book) FetchBooksByAuthorID(authorId string) ([]*model.Book, error) {
	var books []*model.Book
	err := o.DB.Model(&books).Column("book.*").Relation("BookAuthor").
		Where("book_author.author_id = ?", authorId).Select()
	if err != nil {
		return nil, err
	}
	return model.ReshapeBooks(books), nil
}

func (o *Book) FetchBooksByPublisherID(publisherID string) ([]*model.Book, error) {
	var books []*model.Book
	err := o.DB.Model(&books).Column("book.*").Relation("BookPublisher").
		Where("book_publisher.publisher_id = ?", publisherID).Select()
	if err != nil {
		return nil, err
	}
	return model.ReshapeBooks(books), nil
}

func (o *Book) FetchBooksByID(id string) (*model.Book, error) {
	var books []*model.Book
	err := o.DB.Model(&books).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, nil
	}
	return books[0], nil
}
