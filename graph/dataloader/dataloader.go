package dataloader

import (
	"context"
	"github.com/go-pg/pg/v9"
	"net/http"
)

func DataLoader(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		loaders := map[string]interface{}{
			// Put all dataloader in this mapping
			authorLoaderKey:    BuildAuthorLoader(db),
			publisherLoaderKey: BuildPublisherLoader(db),
			topicLoaderKey:     BuildTopicLoader(db),
		}
		ctx := request.Context()
		for k, v := range loaders {
			ctx = context.WithValue(ctx, k, v)
		}
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
