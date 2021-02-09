package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/liujun5885/book_store_gql/middleware"

	"github.com/liujun5885/book_store_gql/graph/model"
)

func (r *rootMutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.RegisterResponse, error) {
	var response = &model.RegisterResponse{}

	if user, _ := r.ORMUser.FetchUserByEmail(&input.Email); user != nil {
		response.Code = model.RegisterCodeFailed
		return response, errors.New(fmt.Sprintf("email %s exists", input.Email))
	}
	if user, _ := r.ORMUser.FetchUserByPhoneNumber(input.PhoneNumber); user != nil {
		response.Code = model.RegisterCodeFailed
		return response, errors.New(fmt.Sprintf("phone %s exists", *input.PhoneNumber))
	}

	user, err := r.ORMUser.CreateUser(
		&input.Email, &input.Password, input.PhoneNumber, input.FirstName, input.LastName,
	)
	if err != nil {
		return nil, err
	}

	authToken, err := user.GenToken()
	if err != nil {
		return nil, err
	}
	response.AuthToken = authToken
	response.Code = model.RegisterCodeSucceeded
	response.User = user

	return response, nil
}

func (r *rootMutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	var response = &model.LoginResponse{}
	user, err := r.ORMUser.Login(input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	authToken, err := user.GenToken()
	if err != nil {
		return nil, err
	}
	response.AuthToken = authToken
	response.Code = model.LoginCodeSucceeded
	response.User = user
	return response, nil
}

func (r *rootQueryResolver) FetchCurrentUser(ctx context.Context) (*model.User, error) {
	user, err := middleware.GetUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
