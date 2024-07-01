package storage

import (
	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/postgres"
)

type Storage interface {
	feedStorage
	memberStorage
	sessionStorage
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
