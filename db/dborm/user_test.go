package dborm

import (
	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserAndFetchUserByEmail(t *testing.T) {
	dbConn := db.Conn(&pg.Options{
		User:     "book_service",
		Database: "ugc",
	})
	user := User{
		DB: dbConn,
	}
	email := "test@test.com"
	password := "password"
	phoneNumber := ""
	firstName := ""
	lastName := ""

	userModel1, err := user.CreateUser(&email, &password, &phoneNumber, &firstName, &lastName)
	if err != nil {
		t.Error(err)
		return
	}
	userModel2, err := user.FetchUserByEmail(&email)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, userModel1.ID, userModel2.ID)
}
