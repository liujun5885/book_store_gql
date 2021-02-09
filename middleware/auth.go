package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/constants"
	"github.com/liujun5885/book_store_gql/db/dborm"
	"github.com/liujun5885/book_store_gql/graph/model"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"strings"
)

var CurrentUserKey = "UserKey"

func AuthMiddleware(db *pg.DB) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			token, err := parseToken(request)
			if err != nil {
				next.ServeHTTP(writer, request)
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				next.ServeHTTP(writer, request)
				return
			}
			userID := claims["jti"].(string)
			userDB := dborm.User{DB: db}
			user, err := userDB.FetchUserByID(&userID)
			if err != nil || user == nil {
				next.ServeHTTP(writer, request)
				return
			}
			fmt.Println(user.ID)
			ctx := context.WithValue(request.Context(), CurrentUserKey, user)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFormToken,
}

func stripBearerPrefixFormToken(token string) (string, error) {
	bearer := "bearer"
	s := strings.SplitN(token, " ", 2)
	if len(s) != 2 || strings.ToLower(s[0]) != bearer {
		return token, nil
	}
	return s[1], nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv(constants.JWTSecret))
		return t, nil
	})
	return jwtToken, errors.Wrap(err, "Parse Token Error")
}

func GetUserFromCTX(ctx context.Context) (*model.User, error) {
	errNoUser := errors.New("no user in context")
	userValue := ctx.Value(CurrentUserKey)
	if userValue == nil {
		return nil, errNoUser
	}
	user, ok := userValue.(*model.User)
	if !ok || user.ID == "" {
		return nil, errNoUser
	}
	return user, nil
}
