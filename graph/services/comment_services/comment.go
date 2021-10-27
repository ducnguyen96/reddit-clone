package comment_services

import (
	"context"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/comment_repository"
)

type CommentService struct {
	repository *comment_repository.CommentRepository
}

func NewCommentService(repository *comment_repository.CommentRepository) *CommentService {
	return &CommentService{repository: repository}
}

func (c *CommentService) Create(ctx context.Context, usr ent.User, input model.CreateCommentInput) (*ent.Comment,error) {
	return c.repository.Create(ctx, usr, input)
}

func (c *CommentService) Query(ctx context.Context, input model.QueryCommentInput) []*ent.Comment {
	return c.repository.QueryComment(ctx, input)
}

func (c *CommentService) GetUserActionStatusForComment(ctx context.Context, commentID uint64, usr *ent.User) (bool, bool) {
	return c.repository.GetUserActionStatusForComment(ctx, commentID, usr)
}

func (c *CommentService) GetOwner(ctx context.Context, comment ent.Comment) *ent.User {
	return c.repository.GetOwner(ctx, comment)
}

func (c *CommentService) GetComment(ctx context.Context, id uint64) *ent.Comment {
	return c.repository.GetComment(ctx, id)
}