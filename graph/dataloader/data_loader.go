package dataloader

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/graph/model"
	"net/http"
	"time"
)

const authorLoaderKey = "authorLoader"
const publisherLoaderKey = "publisherLoader"

func DataLoader(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authorLoader := AuthorLoader{
			maxBatch: 100,
			wait:     1 * time.Second,
			fetch: func(keys []string) ([][]*model.Author, []error) {
				var authors []*model.Author
				err := db.Model(&authors).Column("author.*").
					Relation("BookAuthor").
					Where("book_author.book_id in (?)", pg.In(keys)).
					Select()
				if err != nil {
					return nil, []error{err}
				}
				bookVSAuthor := make(map[string][]*model.Author, len(keys))
				for _, author := range authors {
					bookVSAuthor[author.BookAuthor.BookID] = append(bookVSAuthor[author.BookAuthor.BookID], author)
				}
				result := make([][]*model.Author, len(keys))
				for i, bookID := range keys {
					result[i] = bookVSAuthor[bookID]
				}
				return result, []error{err}
			},
		}

		publisherLoader := PublisherLoader{
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

		ctx := context.WithValue(request.Context(), authorLoaderKey, &authorLoader)
		ctx2 := context.WithValue(ctx, publisherLoaderKey, &publisherLoader)
		next.ServeHTTP(writer, request.WithContext(ctx2))
	})
}

func GetAuthorLoader(ctx context.Context) *AuthorLoader {
	return ctx.Value(authorLoaderKey).(*AuthorLoader)
}

func GetPublisherLoader(ctx context.Context) *PublisherLoader {
	return ctx.Value(publisherLoaderKey).(*PublisherLoader)
}
