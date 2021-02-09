package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/liujun5885/book_store_gql/graph/generated"
	"github.com/liujun5885/book_store_gql/graph/model"
)

func (r *authorResolver) Books(ctx context.Context, obj *model.Author) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

type authorResolver struct{ *Resolver }
