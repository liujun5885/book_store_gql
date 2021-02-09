package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/liujun5885/book_store_gql/graph/generated"
	"github.com/liujun5885/book_store_gql/graph/model"
)

func (r *topicResolver) Books(ctx context.Context, obj *model.Topic) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Topic returns generated.TopicResolver implementation.
func (r *Resolver) Topic() generated.TopicResolver { return &topicResolver{r} }

type topicResolver struct{ *Resolver }
