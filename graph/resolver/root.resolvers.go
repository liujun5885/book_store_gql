package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/liujun5885/book_store_gql/graph/generated"
)

// RootMutation returns generated.RootMutationResolver implementation.
func (r *Resolver) RootMutation() generated.RootMutationResolver { return &rootMutationResolver{r} }

// RootQuery returns generated.RootQueryResolver implementation.
func (r *Resolver) RootQuery() generated.RootQueryResolver { return &rootQueryResolver{r} }

type rootMutationResolver struct{ *Resolver }
type rootQueryResolver struct{ *Resolver }
