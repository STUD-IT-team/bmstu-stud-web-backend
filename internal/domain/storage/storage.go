package storage

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type Storage interface {
	GetAllFeed() ([]domain.Feed, error)
	GetFeed(id int) (domain.Feed, error)
}

type storage struct {
	postgres Postgres
}

func NewStorage(postgres Postgres) *storage {
	return &storage{postgres: postgres}
}

func (s *storage) GetAllFeed() ([]domain.Feed, error) {
	return s.postgres.GetAllFeed()
}

func (s *storage) GetFeed(id int) (domain.Feed, error) {
	return s.postgres.GetFeed(id)
}
