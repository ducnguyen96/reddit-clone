package post_services

import (
	"context"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/post_repository"
)

type PostService struct {
	repository *post_repository.PostRepository
}

func NewPostService(repository *post_repository.PostRepository) *PostService {
	return &PostService{repository: repository}
}

func (p *PostService) Create(ctx context.Context, usr ent.User, input model.CreatePostInput) (*ent.Post,error) {
	return p.repository.Create(ctx, usr, input)
}

func (p *PostService) GetBySlug(ctx context.Context, slug string) *ent.Post {
	return p.repository.FindBySlug(ctx, slug)
}

func (p *PostService) Query(ctx context.Context, input model.QueryPostInput) []*ent.Post {
	return p.repository.QueryPost(ctx, input)
}

func (p *PostService) GetCommunity(ctx context.Context, postId uint64) (*ent.Community,error) {
	return p.repository.GetCommunity(ctx, postId)
}

func (p *PostService) GetOwner(ctx context.Context, postId uint64) (*ent.User,error) {
	return p.repository.GetOwner(ctx, postId)
}

func (p *PostService) GetNumberOfComments(ctx context.Context, postId uint64) int {
	return p.repository.GetNumberOfComments(ctx, postId)
}