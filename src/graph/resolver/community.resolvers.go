package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/ducnguyen96/reddit-clone/graph/generated"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

func (r *mutationResolver) CreateCommunity(ctx context.Context, input model.CreateCommunityInput) (*model.Community, error) {
	usr, err := r.UerService.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	c, err := r.CommunityService.Create(ctx, *usr, input)
	if err != nil {
		return nil, err
	}
	return utils.MapEntGoCommunityToGraphCommunity(c), nil
}

func (r *queryResolver) GetCommunity(ctx context.Context, slug string) (*model.Community, error) {
	c := r.CommunityService.GetBySlug(ctx, slug)
	if c == nil {
		return nil, errors.New("not found")
	}
	return utils.MapEntGoCommunityToGraphCommunity(c), nil
}

func (r *queryResolver) QueryCommunity(ctx context.Context, input model.QueryCommunityInput) (*model.CommunityPagination, error) {
	c := r.CommunityService.Query(ctx, input)
	l := len(c)

	result := make([]*model.Community, l)

	for i, community := range c {
		result[i] = utils.MapEntGoCommunityToGraphCommunity(community)
	}

	return &model.CommunityPagination{
		Length:      l,
		CurrentPage: *input.Page,
		Communities: result,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
