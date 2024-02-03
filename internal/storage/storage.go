package storage

import (
	"context"

	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/postgres"
)

type Storage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	SetSessionCache(id string, value domain.Session)
	FindSessionCache(id string) *domain.Session
	DeleteSessionCache(id string)
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

func (s *storage) GetAllFeed(ctx context.Context) ([]domain.Feed, error) {
	return s.postgres.GetAllFeed(ctx)
}

func (s *storage) GetFeed(ctx context.Context, id int) (domain.Feed, error) {
	return s.postgres.GetFeed(ctx, id)
}

func (s *storage) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.postgres.GetUserByEmail(ctx, email)
}

func (s *storage) SetSessionCache(id string, value domain.Session) {
	s.sessionCache.Put(id, value)
}

func (s *storage) FindSessionCache(id string) *domain.Session {
	return s.sessionCache.Find(id)
}

func (s *storage) DeleteSessionCache(id string) {
	s.sessionCache.Delete(id)
}
