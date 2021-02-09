package dataloader

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/graph/model"
	"time"
)

const authorLoaderKey = "authorLoader"

func BuildAuthorLoader(db *pg.DB) *AuthorLoader {
	return &AuthorLoader{
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
}

func GetAuthorLoader(ctx context.Context) *AuthorLoader {
	return ctx.Value(authorLoaderKey).(*AuthorLoader)
}
