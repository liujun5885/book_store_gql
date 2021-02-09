package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/liujun5885/book_store_gql/graph/generated"
)

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Publisher returns generated.PublisherResolver implementation.
func (r *Resolver) Publisher() generated.PublisherResolver { return &publisherResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Topic returns generated.TopicResolver implementation.
func (r *Resolver) Topic() generated.TopicResolver { return &topicResolver{r} }

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type publisherResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type topicResolver struct{ *Resolver }
