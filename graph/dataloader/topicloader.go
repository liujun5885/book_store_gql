package dataloader

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/graph/model"
	"time"
)

const topicLoaderKey = "topicLoader"

func BuildTopicLoader(db *pg.DB) *TopicLoader {
	return &TopicLoader{
		maxBatch: 100,
		wait:     1 * time.Second,
		fetch: func(keys []string) ([][]*model.Topic, []error) {
			var topics []*model.Topic
			err := db.Model(&topics).Column("topic.*").
				Relation("BookTopic").
				Where("book_topic.book_id in (?)", pg.In(keys)).
				Select()
			if err != nil {
				return nil, []error{err}
			}
			bookVSTopic := make(map[string][]*model.Topic, len(keys))
			for _, topic := range topics {
				bookVSTopic[topic.BookTopic.BookID] = append(bookVSTopic[topic.BookTopic.BookID], topic)
			}
			result := make([][]*model.Topic, len(keys))
			for i, bookID := range keys {
				result[i] = bookVSTopic[bookID]
			}
			return result, []error{err}
		},
	}
}

func GetTopicLoader(ctx context.Context) *TopicLoader {
	return ctx.Value(topicLoaderKey).(*TopicLoader)
}
