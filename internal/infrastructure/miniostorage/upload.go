package miniostorage

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
)

type UploadObject struct {
	BucketName  string
	ObjectName  string
	Data        []byte
	Size        int64
	ContentType string
}

type uploadStorage interface {
	UploadObject(ctx context.Context, u *UploadObject) (string, error)
}

func (s *MinioStorage) UploadObject(ctx context.Context, u *UploadObject) (string, error) {
	reader := bytes.NewReader(u.Data)

	info, err := s.client.PutObject(
		ctx,
		u.BucketName,
		u.ObjectName,
		reader,
		u.Size,
		minio.PutObjectOptions{ContentType: u.ContentType},
	)
	if err != nil {
		return "", err
	}
	return info.Key, nil
}
