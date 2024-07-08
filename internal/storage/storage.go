package storage

import (
	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/postgres"
)

type Storage interface {
	feedStorage
	memberStorage
	guardStorage
	clubStorage
	mediaFileStorage
	eventStorage
}

type storage struct {
	postgres     postgres.Postgres
	sessionCache cache.ICache[string, domain.Session]
}

func NewStorage(postgres postgres.Postgres, sessionCache cache.ICache[string, domain.Session]) *storage {
	return &storage{
		postgres:     postgres,
		sessionCache: sessionCache,
	}
}
