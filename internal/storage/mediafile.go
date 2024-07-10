package storage

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type mediaFileStorage interface {
	GetMediaFile(id int) (*domain.MediaFile, error)
	GetMediaFiles(ids []int) (map[int]domain.MediaFile, error)
}

func (s *storage) GetMediaFile(id int) (*domain.MediaFile, error) {
	return s.postgres.GetMediaFile(id)
}

func (s *storage) GetMediaFiles(ids []int) (map[int]domain.MediaFile, error) {
	return s.postgres.GetMediaFiles(ids)
}
