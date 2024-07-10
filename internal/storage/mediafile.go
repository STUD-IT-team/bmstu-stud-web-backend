package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type mediaFileStorage interface {
	GetMediaFile(id int) (*domain.MediaFile, error)
	GetMediaFiles(ids []int) (map[int]domain.MediaFile, error)
	PutMediaFile(name string, key string) (int, error)
	DeleteMediaFile(id int) error
	GetUnusedMedia(ctx context.Context) ([]domain.MediaFile, error)
	DeleteMediaFiles(ctx context.Context, keys []string) error
}

func (s *storage) GetMediaFile(id int) (*domain.MediaFile, error) {
	return s.postgres.GetMediaFile(id)
}

func (s *storage) GetMediaFiles(ids []int) (map[int]domain.MediaFile, error) {
	return s.postgres.GetMediaFiles(ids)
}

func (s *storage) PutMediaFile(name string, key string) (int, error) {
	return s.postgres.AddMediaFile(name, key)
}

func (s *storage) DeleteMediaFile(id int) error {
	return s.postgres.DeleteMediaFile(id)
}

func (s *storage) GetUnusedMedia(ctx context.Context) ([]domain.MediaFile, error) {
	return s.postgres.GetUnusedMedia(ctx)
}

func (s *storage) DeleteMediaFiles(ctx context.Context, keys []string) error {
	return s.postgres.DeleteMediaFiles(ctx, keys)
}
