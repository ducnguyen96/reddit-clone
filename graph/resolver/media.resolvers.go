package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

func (r *mutationResolver) CreateMedia(ctx context.Context, input model.CreateMediaInput) (*model.Media, error) {
	usr, err := r.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	media, err := r.MediaService.CreateMedia(ctx, input)

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return utils.EntMediaToGraph(media), nil
}
