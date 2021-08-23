// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	Expiration  time.Time `json:"expiration"`
}

type BookPresignObject struct {
	PresignedURL string `json:"presignedUrl"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	Code      LoginCode  `json:"code"`
	User      *User      `json:"user"`
}

type PageCursor struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type PageInfo struct {
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
}

type RegisterInput struct {
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	PhoneNumber *string `json:"phoneNumber"`
	Username    *string `json:"username"`
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
}

type RegisterResponse struct {
	Code      RegisterCode `json:"code"`
	AuthToken *AuthToken   `json:"authToken"`
	User      *User        `json:"user"`
}

type SearchBooksResponse struct {
	PageInfo *PageInfo `json:"pageInfo"`
	Books    []*Book   `json:"books"`
}

type LoginCode string

const (
	LoginCodeSucceeded       LoginCode = "Succeeded"
	LoginCodeInvalidPassword LoginCode = "InvalidPassword"
	LoginCodeInvalidEmail    LoginCode = "InvalidEmail"
)

var AllLoginCode = []LoginCode{
	LoginCodeSucceeded,
	LoginCodeInvalidPassword,
	LoginCodeInvalidEmail,
}

func (e LoginCode) IsValid() bool {
	switch e {
	case LoginCodeSucceeded, LoginCodeInvalidPassword, LoginCodeInvalidEmail:
		return true
	}
	return false
}

func (e LoginCode) String() string {
	return string(e)
}

func (e *LoginCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LoginCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LoginCode", str)
	}
	return nil
}

func (e LoginCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RegisterCode string

const (
	RegisterCodeSucceeded RegisterCode = "Succeeded"
	RegisterCodeFailed    RegisterCode = "Failed"
)

var AllRegisterCode = []RegisterCode{
	RegisterCodeSucceeded,
	RegisterCodeFailed,
}

func (e RegisterCode) IsValid() bool {
	switch e {
	case RegisterCodeSucceeded, RegisterCodeFailed:
		return true
	}
	return false
}

func (e RegisterCode) String() string {
	return string(e)
}

func (e *RegisterCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RegisterCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RegisterCode", str)
	}
	return nil
}

func (e RegisterCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
