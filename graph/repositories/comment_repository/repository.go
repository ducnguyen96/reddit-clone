package comment_repository

import (
	"context"
	"fmt"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/ent/action"
	"github.com/ducnguyen96/reddit-clone/ent/comment"
	"github.com/ducnguyen96/reddit-clone/ent/post"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

type CommentRepository struct {
	readDB  *ent.Client
	writeDB *ent.Client
}

func NewCommentRepository(readDB *ent.Client, writeDB *ent.Client) *CommentRepository {
	return &CommentRepository{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (c *CommentRepository) Create(ctx context.Context, usr ent.User, input model.CreateCommentInput) (*ent.Comment, error) {
	tx, err := c.writeDB.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	po, err := c.readDB.Post.Query().Where(post.ID(utils.StringToUint64(input.PostID))).First(ctx)

	if err != nil {
		return nil, fmt.Errorf("not found post: %w", err)
	}

	var parent *ent.Comment = nil

	if input.ParentID != nil {
		parent = c.readDB.Comment.Query().Where(comment.ID(utils.StringToUint64(*input.ParentID))).FirstX(ctx)
		if parent == nil {
			return nil, fmt.Errorf("not found parent comment")
		}
	}

	coCreate := tx.Comment.Create().
		SetContent(input.Content).
		SetContentMode(utils.GraphContentModeToEntContentMode(input.ContentMode)).
		SetPostID(utils.StringToUint64(input.PostID)).
		AddPosts(po).
		AddUser(&usr)

	if parent != nil {
		coCreate.SetParent(parent)
	}

	co, err := coCreate.Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating new comment: %w", err)
	}

	//if parent != nil {
	//	_, err = tx.Comment.UpdateOne(parent).AddChildren(co).Save(ctx)
	//	if err != nil {
	//		return nil, fmt.Errorf("failed updating parent comment %w", err)
	//	}
	//}

	//_, err = tx.Post.UpdateOne(po).AddComments(co).Save(ctx)
	//
	//if err != nil {
	//	return nil, fmt.Errorf("failed adding comment to post: %w", err)
	//}
	//
	//_, err = tx.User.UpdateOne(&usr).AddComments(co).Save(ctx)
	//
	//if err != nil {
	//	return nil, fmt.Errorf("failed adding comment to user: %w", err)
	//}

	err = tx.Commit()

	if err != nil {
		err := rollback(tx, fmt.Errorf("failed commiting create comment: %w", err))
		return nil, err
	}
	return co, nil
}

func (c *CommentRepository) QueryComment(ctx context.Context, input model.QueryCommentInput) []*ent.Comment {
	limit, page := input.Limit, input.Page
	if limit == nil {
		*limit = 10
	}
	if page == nil {
		*page = 1
	}
	offset := (*page - 1) * *limit

	qr := c.readDB.Comment.Query().
		Where(comment.PostID(utils.StringToUint64(input.PostID))).
		Limit(*limit).
		Offset(offset)

	if input.ParentID != nil {
		qr.Where(comment.HasParentWith(comment.ID(utils.StringToUint64(*input.ParentID))))
	} else {
		qr.Where(comment.Not(comment.HasParent()))
	}

	return qr.
		Order(ent.Desc(comment.FieldUpVotes)).
		Order(ent.Desc(comment.FieldCreatedAt)).
		AllX(ctx)
}

func (c *CommentRepository) GetUserActionStatusForComment(ctx context.Context, commentID uint64, usr *ent.User) (bool, bool) {
	co := c.readDB.User.QueryActions(usr).
		Where(action.TargetTypeEQ(enums.COMMENT)).
		Where(action.Target(commentID)).
		FirstX(ctx)

	if co == nil {
		return false, false
	} else {
		if co.Type == enums.UpVote {
			return true, false
		} else {
			return false, true
		}
	}
}

func (c *CommentRepository) GetOwner(ctx context.Context, comment ent.Comment) *ent.User {
	return c.readDB.Comment.QueryUser(&comment).FirstX(ctx)
}

func (c *CommentRepository) GetComment(ctx context.Context, id uint64) *ent.Comment {
	return c.readDB.Comment.Query().Where(comment.ID(id)).FirstX(ctx)
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
