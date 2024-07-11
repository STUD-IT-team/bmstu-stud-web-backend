package miniostorage

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type listStorage interface {
	GetAllObjectNames(ctx context.Context, bucketName string) ([]string, error)
}

func (s *MinioStorage) GetAllObjectNames(ctx context.Context, bucketName string) ([]string, error) {
	ch := s.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Recursive: true})
	var filenames []string

	for object := range ch {
		if object.Err != nil {
			return nil, object.Err
		}
		if object.Key[len(object.Key)-1] != '/' {
			filenames = append(filenames, object.Key)
		}
	}
	return filenames, nil
}
