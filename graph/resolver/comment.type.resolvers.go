package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ducnguyen96/reddit-clone/graph/generated"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment) ([]*model.Comment, error) {
	usr, _ := r.UserService.GetCurrentUser(ctx)

	limit := 99999
	page := 1

	comments := r.CommentService.Query(ctx, model.QueryCommentInput{
		PostID:   obj.PostID,
		Limit:    &limit,
		Page:     &page,
		ParentID: &obj.ID,
	})

	length := len(comments)

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
	return resultComments, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
