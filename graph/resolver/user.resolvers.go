package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

func (r *mutationResolver) UserCreateAction(ctx context.Context, input model.UserCreateActionInput) (*model.UserAction, error) {
	usr, err := r.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	at, err := r.UserService.CreateAction(ctx, *usr, input)
	if err != nil {
		return nil, err
	}

	return utils.EntUserActionToGraph(at), nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	usr, err := r.UserService.GetCurrentUserNoTokenValid(ctx)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return utils.MapEntGoUserToGraphUser(usr), nil
}
