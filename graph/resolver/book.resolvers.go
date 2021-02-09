package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/liujun5885/book_store_gql/graph/dataloader"
	"github.com/liujun5885/book_store_gql/graph/generated"
	"github.com/liujun5885/book_store_gql/graph/model"
	"github.com/liujun5885/book_store_gql/middleware"
)

func (r *bookResolver) Authors(ctx context.Context, obj *model.Book) ([]*model.Author, error) {
	return dataloader.GetAuthorLoader(ctx).Load(obj.ID)
}

func (r *bookResolver) Publishers(ctx context.Context, obj *model.Book) ([]*model.Publisher, error) {
	return dataloader.GetPublisherLoader(ctx).Load(obj.ID)
}

func (r *bookResolver) Topics(ctx context.Context, obj *model.Book) ([]*model.Topic, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *rootQueryResolver) SearchBooks(ctx context.Context, keyword string) ([]*model.Book, error) {
	if _, err := middleware.GetUserFromCTX(ctx); err != nil {
		return nil, err
	}
	return r.ORMBooks.SearchBooks(keyword)
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

type bookResolver struct{ *Resolver }
