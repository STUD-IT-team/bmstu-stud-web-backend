package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/miniostorage"
	"golang.org/x/crypto/bcrypt"
)

func (s *storage) UploadObject(ctx context.Context, name string, bucketName string, data []byte) (string, error) {
	upl := miniostorage.UploadObject{
		BucketName:  bucketName,
		ObjectName:  name,
		Data:        data,
		Size:        int64(len(data)),
		ContentType: "",
	}
	minioKey, err := s.minio.UploadObject(ctx, &upl)
	if err != nil {
		return "", err
	}
	return minioKey, err
}

const bcryptCost = 10

func (s *storage) UploadObjectBcrypt(ctx context.Context, name string, bucketName string, data []byte) (string, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(name), bcryptCost)
	if err != nil {
		return "", err
	}
	upl := miniostorage.UploadObject{
		BucketName:  bucketName,
		ObjectName:  string(key),
		Data:        data,
		Size:        int64(len(data)),
		ContentType: "",
	}
	minioKey, err := s.minio.UploadObject(ctx, &upl)
	if err != nil {
		return "", err
	}
	return minioKey, err
}

func (s *storage) DeleteObject(ctx context.Context, name string, bucketName string) error {
	req := miniostorage.DeleteObject{
		ObjectName: name,
		BucketName: bucketName,
	}
	return s.minio.DeleteObject(ctx, &req)
}
