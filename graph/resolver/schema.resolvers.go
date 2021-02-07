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

func (r *bookResolver) Authors(ctx context.Context, obj *model.Book) ([]*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *bookResolver) Publishers(ctx context.Context, obj *model.Book) ([]*model.Publisher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.RegisterResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *publisherResolver) Books(ctx context.Context, obj *model.Publisher) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchBooks(ctx context.Context, keyword string) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FetchCurrentUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *topicResolver) Books(ctx context.Context, obj *model.Topic) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Publisher returns generated.PublisherResolver implementation.
func (r *Resolver) Publisher() generated.PublisherResolver { return &publisherResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Topic returns generated.TopicResolver implementation.
func (r *Resolver) Topic() generated.TopicResolver { return &topicResolver{r} }

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type publisherResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type topicResolver struct{ *Resolver }
