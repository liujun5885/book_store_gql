package dborm

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/graph/model"
)

type Book struct {
	DB *pg.DB
}

func (o *Book) SearchBooks(keyword string) ([]*model.Book, error) {
	var books []*model.Book
	err := o.DB.Model(&books).Where("title ILIKE ?", fmt.Sprintf("%%%s%%", keyword)).Select()
	if err != nil {
		println("error here")
		return nil, err
	}

	return model.ReshapeBooks(books), nil
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
