package storage

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type clubStorage interface {
	GetClub(id int) (*domain.Club, error)
}

func (s *storage) GetClub(id int) (*domain.Club, error) {
	return s.postgres.GetClub(id)
}
