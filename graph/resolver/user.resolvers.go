package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/liujun5885/book_store_gql/graph/generated"
	"github.com/liujun5885/book_store_gql/graph/model"
	"github.com/liujun5885/book_store_gql/middleware"
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

func (r *rootMutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	user, err := middleware.GetUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if input.Basic != nil {
		user.PhoneNumber = input.Basic.PhoneNumber
		user.FirstName = input.Basic.FirstName
		user.LastName = input.Basic.LastName
		if _, err := r.ORMUser.UpdateUser(user); err != nil {
			return nil, err
		}
	}

	if input.Profile != nil {
		profile := &model.UserProfile{
			UserID:   user.ID,
			Address:  input.Profile.Address,
			City:     input.Profile.City,
			Province: input.Profile.Province,
			Country:  input.Profile.Country,
			Job:      input.Profile.Job,
			School:   input.Profile.School,
		}
		if _, err := r.ORMUser.UpdateUserProfiles(profile); err != nil {
			return nil, err
		}
	}

	if input.Settings != nil {
		settings := &model.UserSettings{
			UserID:        user.ID,
			KindleAccount: input.Settings.KindleAccount,
		}
		if _, err := r.ORMUser.UpdateUserSettings(settings); err != nil {
			return nil, err
		}
	}
	user.Profile, err = r.ORMUser.FetchProfilesByUserID(&user.ID)
	if err != nil {
		return nil, err
	}
	user.Settings, err = r.ORMUser.FetchSettingsByUserID(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *rootQueryResolver) FetchCurrentUser(ctx context.Context) (*model.User, error) {
	user, err := middleware.GetUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userResolver) Profile(ctx context.Context, obj *model.User) (*model.UserProfile, error) {
	user, err := middleware.GetUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	profile, err := r.ORMUser.FetchProfilesByUserID(&user.ID)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *userResolver) Settings(ctx context.Context, obj *model.User) (*model.UserSettings, error) {
	user, err := middleware.GetUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	settings, err := r.ORMUser.FetchSettingsByUserID(&user.ID)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
