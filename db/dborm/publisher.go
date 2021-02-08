package dborm

import (
	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/graph/model"
)

type Publisher struct {
	DB *pg.DB
}

func (o *Publisher) SearchPublisher(keyword string) ([]*model.Publisher, error) {
	var publishers []*model.Publisher
	err := o.DB.Model(&publishers).Select()
	if err != nil {
		return nil, err
	}

	return publishers, nil
}

func (o *Publisher) FetchPublishersByBookID(bookID string) ([]*model.Publisher, error) {
	var publishers []*model.Publisher
	err := o.DB.Model(&publishers).Column("publisher.*").
		Relation("BookPublisher").
		Where("book_publisher.book_id = ?", bookID).
		Select()
	if err != nil {
		return nil, err
	}

	return publishers, nil
}
