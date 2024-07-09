package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/miniostorage"
)

type mediaFileStorage interface {
	GetMediaFile(id int) (*domain.MediaFile, error)
	GetMediaFiles(ids []int) (map[int]domain.MediaFile, error)
	UploadObject(ctx context.Context, name string, data []byte) (string, error)
}

func (s *storage) GetMediaFile(id int) (*domain.MediaFile, error) {
	return s.postgres.GetMediaFile(id)
}

func (s *storage) GetMediaFiles(ids []int) (map[int]domain.MediaFile, error) {
	return s.postgres.GetMediaFiles(ids)
}

func (s *storage) UploadObject(ctx context.Context, name string, data []byte) (int, error) {
	upl := miniostorage.UploadObject{
		BucketName:  "images",
		ObjectName:  name,
		Data:        data,
		Size:        int64(len(data)),
		ContentType: "",
	}
	key, err := s.minio.UploadObject(ctx, &upl)
	if err != nil {
		return "", err
	}
	id, err := s.postgres.AddMediaFile(name, key)
	return id, err
}
