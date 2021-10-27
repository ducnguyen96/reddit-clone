package post_repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/ent/community"
	"github.com/ducnguyen96/reddit-clone/ent/post"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
	"github.com/gosimple/slug"
)

type PostRepository struct {
	readDB  *ent.Client
	writeDB *ent.Client
}

func NewPostRepository(readDB *ent.Client, writeDB *ent.Client) *PostRepository {
	return &PostRepository{
		readDB:  readDB,
		writeDB: writeDB,
	}
}
func (p *PostRepository) Create(ctx context.Context, usr ent.User, input model.CreatePostInput) (*ent.Post, error) {
	tx, err := p.writeDB.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	co, err := p.readDB.Community.Query().Where(community.ID(utils.StringToUint64(input.CommunityID))).First(ctx)

	if err != nil {
		return nil, err
	}

	sl := slug.Make(input.Title)
	po, err := tx.Post.Create().
		SetTitle(input.Title).
		SetSlug(sl).
		SetContent(input.Content).
		SetType(utils.GraphPostTypeToEntPostType(input.Type)).
		SetContentMode(utils.GraphContentModeToEntContentMode(input.ContentMode)).
		AddOwner(&usr).
		AddCommunity(co).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	err = tx.Commit()

	if err != nil {
		err := rollback(tx, fmt.Errorf("failed commiting create post: %w", err))
		return nil, err
	}
	return po, nil
}

func (p *PostRepository) FindBySlug(ctx context.Context, slug string) *ent.Post {
	return p.readDB.Post.Query().Where(post.Slug(slug)).FirstX(ctx)
}

func (p *PostRepository) QueryPost(ctx context.Context, input model.QueryPostInput) []*ent.Post {
	limit, page := input.Limit, input.Page
	if limit == nil {
		*limit = 10
	}
	if page == nil {
		*page = 1
	}
	offset := (*page - 1) * *limit

	qr := p.readDB.Post.Query().
		Limit(*limit).
		Offset(offset)

	if input.Sort != nil {
		switch *input.Sort {
		case model.SortPostEnumTop:
			qr.Order(ent.Desc(post.FieldCreatedAt)).
				Order(ent.Desc(post.FieldUpVotes))
			break
		case model.SortPostEnumHot:
			qr.Order(func(s *sql.Selector) {
				s.OrderExpr(sql.ExprFunc(func(b *sql.Builder) {
					b.WriteString("").Ident(post.FieldUpVotes).WriteOp(sql.OpAdd).Ident(post.FieldDownVotes).WriteString("DESC")
				}))
			}).Unique(false).
				Order(ent.Desc(post.FieldID))
			break
		case model.SortPostEnumNew:
			qr.Order(ent.Desc(post.FieldCreatedAt))
			break
		default:
			qr.Order(func(s *sql.Selector) {
				s.OrderExpr(sql.ExprFunc(func(b *sql.Builder) {
					b.WriteString("").Ident(post.FieldUpVotes).WriteOp(sql.OpSub).Ident(post.FieldDownVotes).WriteString("DESC")
				}))
			}).
				Unique(false).
				Order(ent.Desc(post.FieldID))
		}
	}

	return qr.
		AllX(ctx)
}

func (p *PostRepository) GetNumberOfComments(ctx context.Context, postId uint64) int {
	po := p.readDB.Post.Query().Where(post.ID(postId)).FirstX(ctx)
	return p.readDB.Post.QueryComments(po).CountX(ctx)
}

func (p *PostRepository) GetCommunity(ctx context.Context, postId uint64) (*ent.Community, error) {
	po := p.readDB.Post.Query().Where(post.ID(postId)).FirstX(ctx)
	return p.readDB.Post.QueryCommunity(po).First(ctx)
}

func (p *PostRepository) GetOwner(ctx context.Context, postId uint64) (*ent.User, error) {
	po := p.readDB.Post.Query().Where(post.ID(postId)).FirstX(ctx)
	return p.readDB.Post.QueryOwner(po).First(ctx)
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
