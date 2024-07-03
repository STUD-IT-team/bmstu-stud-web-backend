package storage

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type clubStorage interface {
	GetClub(id int) (*domain.Club, error)
	GetAllClub() ([]domain.Club, error)
	GetClubOrgs(clubID int) ([]domain.ClubOrg, error)
	GetClubSubOrgs(clubID int) ([]domain.ClubOrg, error)
	GetAllClubOrgs() ([]domain.ClubOrg, error)
	AddClub(c *domain.Club) (int, error)
	AddOrgs(orgs []domain.ClubOrg) error
}

func (s *storage) GetClub(id int) (*domain.Club, error) {
	return s.postgres.GetClub(id)
}

func (s *storage) GetAllClub() ([]domain.Club, error) {
	return s.postgres.GetAllClub()
}

func (s *storage) GetClubOrgs(clubID int) ([]domain.ClubOrg, error) {
	return s.postgres.GetClubOrgs(clubID)
}

func (s *storage) GetClubSubOrgs(clubID int) ([]domain.ClubOrg, error) {
	return s.postgres.GetClubSubOrgs(clubID)
}

func (s *storage) GetAllClubOrgs() ([]domain.ClubOrg, error) {
	return s.postgres.GetAllClubOrgs()
}

func (s *storage) AddClub(c *domain.Club) (int, error) {
	return s.postgres.AddClub(c)
}

func (s *storage) AddOrgs(orgs []domain.ClubOrg) error {
	return s.postgres.AddOrgs(orgs)
}
