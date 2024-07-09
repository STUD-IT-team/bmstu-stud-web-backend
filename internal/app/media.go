package app

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
)

type mediaStorage interface {
	UploadObject(ctx context.Context, name string, data []byte) (int, error)
}

type MediaService struct {
	storage mediaStorage
}

func NewMediaService(storage mediaStorage) *MediaService {
	return &MediaService{storage: storage}
}

func (s *MediaService) PostObject(ctx context.Context, req *requests.PostMedia) (int, error) {
	id, err := s.storage.UploadObject(ctx, req.Name, req.Data)
	return id, err
}
