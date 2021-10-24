package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ducnguyen96/reddit-clone/graph/generated"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

func (r *postResolver) Community(ctx context.Context, obj *model.Post) (*model.Community, error) {
	co, err := r.PostService.GetCommunity(ctx, utils.StringToUint64(obj.ID))
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	return utils.MapEntGoCommunityToGraphCommunity(co), nil
}

func (r *postResolver) Owner(ctx context.Context, obj *model.Post) (*model.User, error) {
	usr, err := r.PostService.GetOwner(ctx, utils.StringToUint64(obj.ID))
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	return utils.MapEntGoUserToGraphUser(usr), nil
}

func (r *postResolver) NumberOfComments(ctx context.Context, obj *model.Post) (int, error) {
	return r.PostService.GetNumberOfComments(ctx, utils.StringToUint64(obj.ID)), nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
