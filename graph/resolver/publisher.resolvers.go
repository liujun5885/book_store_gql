package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/liujun5885/book_store_gql/graph/generated"
	"github.com/liujun5885/book_store_gql/graph/model"
)

func (r *publisherResolver) Books(ctx context.Context, obj *model.Publisher) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Publisher returns generated.PublisherResolver implementation.
func (r *Resolver) Publisher() generated.PublisherResolver { return &publisherResolver{r} }

type publisherResolver struct{ *Resolver }
