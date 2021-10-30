package media_repository

import (
	"context"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

type MediaRepository struct {
	readDB  *ent.Client
	writeDB *ent.Client
}

func NewMediaRepository(readDB *ent.Client, writeDB *ent.Client) *MediaRepository {
	return &MediaRepository{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (m *MediaRepository) CreateMedia(ctx context.Context, input model.CreateMediaInput) (*ent.Media, error) {
	return m.writeDB.Media.Create().
		SetURL(input.URL).
		SetType(utils.GraphMediaTypeToEntMediaType(input.Type)).Save(ctx)
}