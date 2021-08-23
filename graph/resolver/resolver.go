package resolver

import "github.com/liujun5885/book_store_gql/db/dborm"
import "github.com/liujun5885/book_store_gql/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ORMBooks     dborm.Book
	ORMAuthor    dborm.Author
	ORMPublisher dborm.Publisher
	ORMUser      dborm.User
	ORMTopic     dborm.Topic
	S3Client     *services.S3Client
}
