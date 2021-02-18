package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/liujun5885/book_store_gql/middleware"

	"github.com/liujun5885/book_store_gql/graph/dataloader"
	"github.com/liujun5885/book_store_gql/graph/generated"
	"github.com/liujun5885/book_store_gql/graph/model"
)

func (r *bookResolver) Authors(ctx context.Context, obj *model.Book) ([]*model.Author, error) {
	return dataloader.GetAuthorLoader(ctx).Load(obj.ID)
}

func (r *bookResolver) Publishers(ctx context.Context, obj *model.Book) ([]*model.Publisher, error) {
	return dataloader.GetPublisherLoader(ctx).Load(obj.ID)
}

func (r *bookResolver) Topics(ctx context.Context, obj *model.Book) ([]*model.Topic, error) {
	return dataloader.GetTopicLoader(ctx).Load(obj.ID)
}

func (r *rootQueryResolver) SearchBooks(ctx context.Context, keyword string, pageCursor model.PageCursor) (*model.SearchBooksResponse, error) {
	if _, err := middleware.GetUserFromCTX(ctx); err != nil {
		return nil, err
	}

	if len(keyword) < 2 || len(keyword) > 128 {
		return nil, errors.New("the length of keywords should be more than 1 and less than 128")
	}
	if pageCursor.PageSize <= 0 || pageCursor.PageSize > 100 || pageCursor.Page < 0 {
		return nil, errors.New("PageSize should be more than 0 and be less than 100, Page cannot be negative")
	}

	return r.ORMBooks.SearchBooks(keyword, pageCursor)
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

type bookResolver struct{ *Resolver }
