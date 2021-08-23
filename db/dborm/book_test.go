package dborm

import (
	"github.com/go-pg/pg/v9"
	"github.com/liujun5885/book_store_gql/db"
	"testing"
)

func TestFetchBookByID(t *testing.T) {
	dbConn := db.Conn(&pg.Options{
		User:     "book_store",
		Database: "book_store_assets",
		Password: "aaaaa",
		Addr:     "localhost:5432",
	})
	book := Book{
		DB: dbConn,
	}
	bookID := "397faa60-80eb-470d-8101-13eb2980d16c"
	bookObj, err := book.FetchBooksByID(bookID)
	if err != nil {
		t.Error(err)
	}
	if bookObj == nil {
		t.Errorf("book %s doesn't exist", bookID)
	}
}
