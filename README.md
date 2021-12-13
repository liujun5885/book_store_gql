# book_store_gql

## Prerequisite

- go-pg package: https://pg.uptrace.dev/
- gqlgen, the tool we are using to generate graphql server code: https://gqlgen.com/getting-started/

## Init repo. We don't need to run it. Because we already ran it when we init this repo

```bash
go run github.com/99designs/gqlgen init
```

## How to add your graphQL Model

### Add model in file `graph/book.graphql`

```graphql
extend type RootQuery {
    searchBooks(keyword: String!, pageCursor: PageCursor!): SearchBooksResponse
    generateBookPresignObject(id: String!):BookPresignObject
}

type SearchBooksResponse {
    pageInfo: PageInfo!
    books: [Book]
}

type BookPresignObject {
    presignedUrl: String!
}

type Book {
    id: ID!
    title: String!
    description: String!
    descriptionTrimmed: String!
    pages: Int!
    language: String!
    rating: Int!
    reviews: Int!
    authors: [Author]
    publishers: [Publisher]
    topics: [Topic]
    coverURL: String!
    url: String!
    issuedAt: Time!
    createdAt: Time!
    updatedAt: Time!
    type: String!
}

```

### run command to generate code

```bash
go generate ./...
```

### all generated fill will go to folder `graph`

### all `books` resolvers will goto `resolver/book.resolvers.go`. You need to implement it yourself

#### for example:

```
func (r *rootQueryResolver) SearchBooks(ctx context.Context, keyword string, pageCursor model.PageCursor) (*model.SearchBooksResponse, error) {
    // do something here.
}
```

### `gqlgen` will also create data model for you. If you want to customize it, cut the definition from model `graph/model/models_gen.go` and put it into `graph/model/models_custom.go`. For example:

```go
// original model
package model

// ------------ in file model_gen.go
// original model
type Author struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Nationality string    `json:"nationality"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Books       []*Book   `json:"books"`
}

// ----------- in file model_custom.go
// modified model
type Author struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Nationality string    `json:"nationality"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	// customized fields
	//Books       []*Book    `json:"books"`
	BookAuthor *BookAuthor `json:"-"`
}
```

> you can also put modified model in other files, but should be under package `model`. for example, put it into `auth.model.go`

## add dataloader

### run following command to generate code

```bash
go run github.com/vektah/dataloaden AuthorLoader string "[]*github.com/liujun5885/book_store_gql/graph/model.Author"
```

### it will generate the file `authorloader_gen.go`

### put file `authorloader_gen.go` in folder `graph/dataloader`

### In the future, if you want to generate any dataloader please also put into folder `graph/dataloader`

### in file `graph/dataloader/dataloader.go`, it's the implementation of querying data from DB, for example:

```go
package dataloader

func DataLoader(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		loaders := map[string]interface{}{
			// Put all dataloader in this mapping
			authorLoaderKey:    BuildAuthorLoader(db),
			publisherLoaderKey: BuildPublisherLoader(db),
		}
		ctx := request.Context()
		for k, v := range loaders {
			ctx = context.WithValue(ctx, k, v)
		}
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
```

## referring to

### https://dev.to/glyphack/introduction-to-graphql-server-with-golang-1npk

### https://github.com/99designs/gqlgen/issues/785#issue-465696123

# add sqitch for project

## 1. create folder in project

referring to: [Sqitch Tutorial](https://sqitch.org/docs/manual/sqitchtutorial/)

```shell
mkdir -p sqitch/book_store_ugc
```

## 2. create user for project

``` postgresql
CREATE USER book_store PASSWORD 'aaaaa';
CREATE DATABASE book_store_ugc;
CREATE DATABASE book_store_assets;
GRANT ALL on database book_store_ugc to book_store;
GRANT ALL on database book_store_assets to book_store;
```

## 3. init sqitch project

```shell
sqitch init book_store_ugc --uri  https://github.com/liujun5885/book_store_gql/ --engine pg
```

## 4. add sqitch target. I will use local pg for dev, so I created a `local` target.

```shell
sqitch target add local db:pg://book_store:aaaaa@localhost:5432/book_store_ugc
```

## 5. add `local` target to engine `pg`.

```shell
sqitch engine add pg --target local 
```

## 6. test it under folder `sqitch/book_store_ugc`

```shell
sqitch deploy
sqitch verify
sqitch revert
```