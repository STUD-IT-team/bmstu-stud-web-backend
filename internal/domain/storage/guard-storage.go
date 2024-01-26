package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/postgres"
)

type GuardStorage interface {
	GetUserID(ctx context.Context, email string, password string) (userID string, err error)
}

type guardStorage struct {
	postgres postgres.Postgres
}

func NewGuardStorage(postgres postgres.Postgres) *guardStorage {
	return &guardStorage{postgres: postgres}
}

func (s *guardStorage) GetUserID(ctx context.Context, email string, password string) (userID string, err error) {
	userID, err = s.postgres.GetUserID(ctx, email, password)
	return
}
