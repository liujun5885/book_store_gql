package dborm

import (
	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/graph/model"
)

type Author struct {
	DB *pg.DB
}

func (o *Author) SearchAuthors(keyword string) ([]*model.Author, error) {
	var authors []*model.Author
	err := o.DB.Model(&authors).Select()
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (o *Author) FetchAuthorsByBookID(bookID string) ([]*model.Author, error) {
	var authors []*model.Author
	err := o.DB.Model(&authors).Column("author.*").
		Relation("BookAuthor").
		Where("book_author.book_id = ?", bookID).
		Select()
	if err != nil {
		return nil, err
	}
	return authors, nil
}
