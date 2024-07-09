package miniostorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ObjectStorage interface {
	uploadStorage
	downloadStorage
}

type MinioStorage struct {
	client *minio.Client
}

func NewMinioStorage(minioURL string, minioUser string, minioPassword string, ssl bool) (ObjectStorage, error) {
	var err error
	client, err := minio.New(minioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(minioUser, minioPassword, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	return MinioStorage{client: client}, nil
}
