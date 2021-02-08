package dborm

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/graph/model"
	"strings"
)

type User struct {
	DB *pg.DB
}

func (u *User) FetchUserByID(userID *string) (*model.User, error) {
	if userID == nil || *userID == "" {
		return nil, nil
	}
	user := model.User{}
	err := u.DB.Model(&user).Where("id = ?", *userID).Select()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FetchUserByEmail(email *string) (*model.User, error) {
	if email == nil || *email == "" {
		return nil, nil
	}
	user := model.User{}
	err := u.DB.Model(&user).Where("email = ?", strings.ToLower(*email)).Select()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FetchUserByPhoneNumber(phoneNumber *string) (*model.User, error) {
	if phoneNumber == nil || *phoneNumber == "" {
		return nil, nil
	}
	user := model.User{}
	err := u.DB.Model(&user).Where("phone_number = ?", strings.ToLower(*phoneNumber)).Select()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) CreateUser(email, password, phoneNumber, firstName, lastName *string) (*model.User, error) {
	if email == nil || *email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == nil || *password == "" {
		return nil, errors.New("password cannot be empty")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	user := &model.User{
		ID:       uuid.New().String(),
		Email:    *email,
		Password: string(hashedPassword),
	}
	if phoneNumber != nil {
		user.PhoneNumber = *phoneNumber
	}
	if firstName != nil {
		user.FirstName = *firstName
	}
	if lastName != nil {
		user.LastName = *lastName
	}
	_, err := u.DB.Model(user).Insert()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Login(email, password string) (*model.User, error) {
	user := &model.User{}
	if err := u.DB.Model(user).Where("email = ?", email).Select(); err != nil {
		return nil, err
	}
	if user.ID == "" {
		return nil, errors.New(fmt.Sprintf("user %s does not exist", email))
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New(fmt.Sprintf("invalid password"))
	}
	return user, nil
}
