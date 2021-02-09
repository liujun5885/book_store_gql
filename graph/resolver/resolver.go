package resolver

import (
	"context"
	"fmt"
	"github.com/liujun5885/book_store_gql/db/dborm"
	"github.com/liujun5885/book_store_gql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ORMBooks     dborm.Book
	ORMAuthor    dborm.Author
	ORMPublisher dborm.Publisher
	ORMUser      dborm.User
	ORMTopic     dborm.Topic
}

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
