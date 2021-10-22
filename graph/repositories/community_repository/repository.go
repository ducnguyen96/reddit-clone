package community_repository

import (
	"context"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/ent/community"
	"github.com/ducnguyen96/reddit-clone/ent/user"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
	"github.com/gosimple/slug"
)

type CommunityRepository struct {
	readDB  *ent.Client
	writeDB *ent.Client
}

func NewCommunityRepository(readDB *ent.Client, writeDB *ent.Client) *CommunityRepository {
	return &CommunityRepository{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (r *CommunityRepository) Create(ctx context.Context, user ent.User, input model.CreateCommunityInput) (*ent.Community, error) {
	sl := slug.Make(input.Name)
	communityCreate, err := r.writeDB.Community.Create().
		SetName(input.Name).
		SetSlug(sl).
		SetType(utils.GraphCommunityTypeToCommunityType(input.Type)).
		SetIsAdult(input.IsAdult).
		AddAdmins(&user).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return communityCreate, nil
}

func (r *CommunityRepository) FindById(ctx context.Context, id uint64) *ent.Community {
	return r.readDB.Community.Query().Where(community.ID(id)).FirstX(ctx)
}

func (r *CommunityRepository) FindBySlug(ctx context.Context, slug string) *ent.Community {
	return r.readDB.Community.Query().Where(community.Slug(slug)).FirstX(ctx)
}

func (r *CommunityRepository) QueryCommunity(ctx context.Context, input model.QueryCommunityInput, usr *ent.User) []*ent.Community {
	query := r.readDB.Community.Query()
	if usr != nil {
		query.Where(community.HasAdminsWith(user.ID(usr.ID)))
	}

	limit, page := input.Limit, input.Page

	if limit == nil {
		*limit = 10
	}
	if page == nil {
		*page = 1
	}
	offset := (*page-1)**limit
	return query.
		Limit(*limit).
		Offset(offset).
		Order(ent.Asc(community.FieldCreatedAt)).
		AllX(ctx)
}

func (r *CommunityRepository) NumberOfMembers(ctx context.Context, id uint64) int {
	return r.readDB.Community.Query().Where(community.ID(id)).CountX(ctx)
}