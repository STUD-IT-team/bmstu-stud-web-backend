package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/postgres"
)

type Storage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}

type storage struct {
	postgres postgres.Postgres
}

func NewStorage(postgres postgres.Postgres) *storage {
	return &storage{postgres: postgres}
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
