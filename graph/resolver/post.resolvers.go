package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*model.Post, error) {
	usr, err := r.UerService.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	po, err := r.PostService.Create(ctx, *usr, input)
	if err != nil {
		return nil, err
	}
	return utils.EntGoPostToGraphPost(po), nil
}

func (r *queryResolver) GetPost(ctx context.Context, slug string) (*model.Post, error) {
	c := r.PostService.GetBySlug(ctx, slug)
	if c == nil {
		return nil, errors.New("not found")
	}
	return utils.EntGoPostToGraphPost(c), nil
}

func (r *queryResolver) QueryPost(ctx context.Context, input model.QueryPostInput) (*model.PostPagination, error) {
	c := r.PostService.Query(ctx, input)

	l := len(c)

	result := make([]*model.Post, l)

	for i, post := range c {
		result[i] = utils.EntGoPostToGraphPost(post)
	}

	return &model.PostPagination{
		Length:      l,
		CurrentPage: *input.Page,
		Posts:       result,
	}, nil
}
