package media_services

import (
	"context"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/media_repository"
)

type MediaService struct {
	repository *media_repository.MediaRepository
}

func NewMediaService(repository *media_repository.MediaRepository) *MediaService {
	return &MediaService{repository: repository}
}

func (m *MediaService) CreateMedia(ctx context.Context, input model.CreateMediaInput) (*ent.Media, error) {
	return m.repository.CreateMedia(ctx, input)
}