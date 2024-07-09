package miniostorage

import "context"

type DownloadObject struct {
	BucketName  string
	ObjectName  string
	Data        []byte
	Size        int64
	ContentType string
}

type DownloadRequest struct {
	BucketName string
	ObjectName string
}

type downloadStorage interface {
	DownloadObject(ctx context.Context, req DownloadRequest) (DownloadObject, error)
}
