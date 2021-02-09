package dataloader

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/graph/model"
	"time"
)

const publisherLoaderKey = "publisherLoader"

func BuildPublisherLoader(db *pg.DB) *PublisherLoader {
	return &PublisherLoader{
		maxBatch: 100,
		wait:     1 * time.Second,
		fetch: func(keys []string) ([][]*model.Publisher, []error) {
			var publishers []*model.Publisher
			err := db.Model(&publishers).Column("publisher.*").
				Relation("BookPublisher").
				Where("book_publisher.book_id in (?)", pg.In(keys)).
				Select()
			if err != nil {
				return nil, []error{err}
			}
			bookVSPublisher := make(map[string][]*model.Publisher, len(keys))
			for _, publisher := range publishers {
				bookVSPublisher[publisher.BookPublisher.BookID] = append(bookVSPublisher[publisher.BookPublisher.BookID], publisher)
			}
			result := make([][]*model.Publisher, len(keys))
			for i, bookID := range keys {
				result[i] = bookVSPublisher[bookID]
			}
			return result, []error{err}
		},
	}
}

func GetPublisherLoader(ctx context.Context) *PublisherLoader {
	return ctx.Value(publisherLoaderKey).(*PublisherLoader)
}
