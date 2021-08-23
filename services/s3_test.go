package services

import (
	"fmt"
	"github.com/liujun5885/book_store_gql/constants"
	"testing"
)

func TestGetS3KeyByBookID(t *testing.T) {
	s3, err := NewS3Session(
		"us-east-1", "all-ebook-bucket", constants.S3AccessKey, constants.S3AssessSecret,
	)
	if err != nil {
		t.Errorf("failed to create session")
		return
	}

	bookId := "01120100011SI"
	key, err := s3.GetS3KeyByBookID(bookId)
	if err != nil {
		t.Errorf("failed to get key for book %s!, error: %v", bookId, err)
		return
	}

	if key == "" {
		t.Errorf("book: %s doesn't exist", bookId)
		return
	}
}

func TestGetS3KeyByBookIDNotExist(t *testing.T) {
	s3, err := NewS3Session(
		"us-east-1", "all-ebook-bucket", constants.S3AccessKey, constants.S3AssessSecret,
	)
	if err != nil {
		t.Errorf("failed to create session")
		return
	}

	bookId := "invalid-book-id"
	key, err := s3.GetS3KeyByBookID(bookId)
	if err != nil {
		t.Errorf("failed to get key for book %s!, error: %v", bookId, err)
		return
	}

	if key != "" {
		t.Errorf("key of book: %s is empty, but it's %s", bookId, key)
		return
	}
}

func TestNewSignedGetURL(t *testing.T) {
	s3, err := NewS3Session(
		"us-east-1", "all-ebook-bucket", constants.S3AccessKey, constants.S3AssessSecret,
	)
	if err != nil {
		t.Errorf("failed to create session")
		return
	}

	bookId := "0028636120"
	url, err := s3.NewSignedGetURL(bookId)
	if err != nil {
		t.Errorf("failed to get key for book %s!, error: %v", bookId, err)
		return
	}
	fmt.Printf("url: %s\n", url)
	if url == "" {
		t.Errorf("book: %s doesn't exist", bookId)
		return
	}
}
