package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/liujun5885/book_store_gql/graph/model"
)

func (r *rootMutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.RegisterResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *rootMutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *rootQueryResolver) FetchCurrentUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
