package miniostorage

import "context"

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
