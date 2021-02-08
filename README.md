# book_store_gql

## Init repo. We don't need to run it. Because we already ran it when we init this repo
```bash
go run github.com/99designs/gqlgen init
```

## How to add your graphQL Model
### Add model in file `graph/schema.graphqls`
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