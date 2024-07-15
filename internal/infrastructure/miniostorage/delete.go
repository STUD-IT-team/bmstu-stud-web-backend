package miniostorage

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type DeleteObject struct {
	BucketName string
	ObjectName string
}

type deleteStorage interface {
	DeleteObject(ctx context.Context, u *DeleteObject) error
}

func (s *MinioStorage) DeleteObject(ctx context.Context, u *DeleteObject) error {
	err := s.client.RemoveObject(
		ctx,
		u.BucketName,
		u.ObjectName,
		minio.RemoveObjectOptions{},
	)
	return err
}
