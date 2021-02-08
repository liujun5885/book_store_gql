package dborm

import (
	"github.com/go-pg/pg/v9"
	"testing"

	"github.com/liujun5885/book_store_gql/db"
)

func TestFetchAuthorByBookID(t *testing.T) {
	dbConn := db.Conn(&pg.Options{
		User:     "book_service",
		Database: "assets",
	})
	author := Author{
		DB: dbConn,
	}
	_, err := author.FetchAuthorsByBookID("899423b5-3f03-4a18-8bab-7370ca2d93aa")
	if err == nil {
		t.Error(err)
	}
}
