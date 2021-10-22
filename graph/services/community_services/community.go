package community_services

import (
	"context"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/community_repository"
)

type CommunityService struct {
	repository *community_repository.CommunityRepository
}

func NewCommunityService(repository *community_repository.CommunityRepository) *CommunityService {
	return &CommunityService{repository: repository}
}

func (c *CommunityService) Create(ctx context.Context, usr ent.User, input model.CreateCommunityInput) (*ent.Community, error) {
	return c.repository.Create(ctx, usr, input)
}

func (c *CommunityService) GetBySlug(ctx context.Context, slug string) *ent.Community {
	return c.repository.FindBySlug(ctx, slug)
}

func (c *CommunityService) Query(ctx context.Context, input model.QueryCommunityInput, usr *ent.User) []*ent.Community {
	return c.repository.QueryCommunity(ctx, input, usr)
}

func (c *CommunityService) NumberOfMembers(ctx context.Context, id uint64) int {
	return c.repository.NumberOfMembers(ctx, id)
}
