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

func (r *mutationResolver) CreateComment(ctx context.Context, input model.CreateCommentInput) (*model.Comment, error) {
	usr, err := r.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("user not found %w", err)
	}

	co, err := r.CommentService.Create(ctx, *usr, input)
	if err != nil {
		return nil, fmt.Errorf("failed creating comment %w", err)
	}

	mapped := utils.EntCommentToGraph(co, false, false)
	mapped.Owner = utils.MapEntGoUserToGraphUser(usr)
	mapped.Replies = []*model.Comment{}

	return mapped, nil
}

func (r *queryResolver) QueryComment(ctx context.Context, input model.QueryCommentInput) (*model.CommentPagination, error) {
	usr, _ := r.UserService.GetCurrentUser(ctx)

	comments := r.CommentService.Query(ctx, input)
	length := len(comments)

	var page int
	if input.Page == nil {
		page = 1
	} else {
		page = *input.Page
	}

	resultComments := make([]*model.Comment, length)

	for i, comment := range comments {
		if usr == nil {
			resultComments[i] = utils.EntCommentToGraph(comment, false, false)
		} else {
			isUpVoted, isDownVoted := r.CommentService.GetUserActionStatusForComment(ctx, comment.ID, usr)
			resultComments[i] = utils.EntCommentToGraph(comment, isUpVoted, isDownVoted)
		}
		resultComments[i].Owner = utils.MapEntGoUserToGraphUser(r.CommentService.GetOwner(ctx, *comment))
	}

	return &model.CommentPagination{
		Length:      length,
		CurrentPage: page,
		Comments:    resultComments,
	}, nil
}

func (r *queryResolver) GetComment(ctx context.Context, id string) (*model.Comment, error) {
	usr, _ := r.UserService.GetCurrentUser(ctx)
	co := r.CommentService.GetComment(ctx, utils.StringToUint64(id))

	if usr == nil {
		mapped := utils.EntCommentToGraph(co, false, false)
		mapped.Owner = utils.MapEntGoUserToGraphUser(r.CommentService.GetOwner(ctx, *co))
		return mapped, nil
	}
	isUpVoted, isDownVoted := r.CommentService.GetUserActionStatusForComment(ctx, co.ID, usr)
	mapped := utils.EntCommentToGraph(co, isUpVoted, isDownVoted)
	mapped.Owner = utils.MapEntGoUserToGraphUser(r.CommentService.GetOwner(ctx, *co))
	return mapped, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
