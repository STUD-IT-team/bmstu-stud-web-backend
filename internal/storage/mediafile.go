package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/miniostorage"
	"golang.org/x/crypto/bcrypt"
)

type mediaFileStorage interface {
	GetMediaFile(id int) (*domain.MediaFile, error)
	GetMediaFiles(ids []int) (map[int]domain.MediaFile, error)
	UploadObject(ctx context.Context, name string, data []byte) (int, error)
	UploadObjectBcrypt(ctx context.Context, name string, data []byte) (int, error)
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
		return 0, err
	}
	id, err := s.postgres.AddMediaFile(name, key)
	return id, err
}

const bcryptCost = 10

func (s *storage) UploadObjectBcrypt(ctx context.Context, name string, data []byte) (int, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(name), bcryptCost)
	if err != nil {
		return 0, err
	}
	upl := miniostorage.UploadObject{
		BucketName:  "images",
		ObjectName:  string(key),
		Data:        data,
		Size:        int64(len(data)),
		ContentType: "",
	}
	_, err = s.minio.UploadObject(ctx, &upl)
	if err != nil {
		return 0, err
	}
	id, err := s.postgres.AddMediaFile(name, string(key))
	return id, err
}
