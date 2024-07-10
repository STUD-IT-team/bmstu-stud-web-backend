package storage

import (
	// "github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/miniostorage"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/postgres"
)

type Storage interface {
	feedStorage
	memberStorage
	guardStorage
	clubStorage
	mediaFileStorage
	eventStorage
	minioStorage
	documentsStorage
}

type storage struct {
	postgres     postgres.Postgres
	sessionCache cache.SessionCache
	minio        miniostorage.ObjectStorage
}

func NewStorage(postgres postgres.Postgres, sessionCache cache.SessionCache, minio miniostorage.ObjectStorage) *storage {
	return &storage{
		postgres:     postgres,
		sessionCache: sessionCache,
		minio:        minio,
	}
}
