# book_store_gql

## Init repo. We don't need to run it. Because we already ran it when we init this repo

```bash
go run github.com/99designs/gqlgen init
```

## How to add your graphQL Model

### Add model in file `graph/topic.graphql`

```graphql
type Topic {
    id: ID!
    name: String!
    slug: String!
    score: Int
    createdAt: Time!
    updatedAt: Time!
    books: [Book]
}
```

### run command to generate code

```bash
go generate ./...
```

### all generated fill will go to folder `graph`

### all `topic` resolvers will goto `topic.resolvers.go`. You need to implement it yourself

#### for example:

```
func (r *topicResolver) Books(ctx context.Context, obj *model.Topic) ([]*model.Book, error) {
// write your code here
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