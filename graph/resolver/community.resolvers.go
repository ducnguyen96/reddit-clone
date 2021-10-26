package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

func (r *mutationResolver) CreateCommunity(ctx context.Context, input model.CreateCommunityInput) (*model.Community, error) {
	usr, err := r.UserService.GetCurrentUser(ctx)
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
	var c []*ent.Community
	if input.OnlyMine != nil && *input.OnlyMine == true {
		usr, err := r.UserService.GetCurrentUser(ctx)
		if err != nil {
			c = []*ent.Community{}
		} else {
			c = r.CommunityService.Query(ctx, input, usr)
		}
	} else {
		c = r.CommunityService.Query(ctx, input, nil)
	}
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

func (r *queryResolver) IsCommunityNameExisted(ctx context.Context, name string) (bool, error) {
	c := r.CommunityService.GetBySlug(ctx, name)
	return c != nil, nil
}
