package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ducnguyen96/reddit-clone/graph/generated"
	"github.com/ducnguyen96/reddit-clone/graph/model"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	usr, err := r.UerService.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
