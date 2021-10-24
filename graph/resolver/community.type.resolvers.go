package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ducnguyen96/reddit-clone/graph/generated"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

func (r *communityResolver) NumberOfMember(ctx context.Context, obj *model.Community) (int, error) {
	id := utils.StringToUint64(obj.ID)
	return r.CommunityService.NumberOfMembers(ctx, id), nil
}

// Community returns generated.CommunityResolver implementation.
func (r *Resolver) Community() generated.CommunityResolver { return &communityResolver{r} }

type communityResolver struct{ *Resolver }
