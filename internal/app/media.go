package app

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
)

type mediaStorage interface {
	UploadObject(ctx context.Context, name string, data []byte) (int, error)
	UploadObjectBcrypt(ctx context.Context, name string, data []byte) (int, error)
}

type MediaService struct {
	storage mediaStorage
}

func NewMediaService(storage mediaStorage) *MediaService {
	return &MediaService{storage: storage}
}

func (s *MediaService) PostObject(ctx context.Context, req *requests.PostMedia) (int, error) {
	id, err := s.storage.UploadObject(ctx, req.Name, req.Data)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (s *MediaService) PostObjectBcrypt(ctx context.Context, req *requests.PostMedia) (int, error) {
	id, err := s.storage.UploadObjectBcrypt(ctx, req.Name, req.Data)
	if err != nil {
		return 0, err
	}
	return id, err
}
