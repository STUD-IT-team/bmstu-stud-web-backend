package storage

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type clubStorage interface {
	GetClub(id int) (*domain.Club, error)
	GetAllClub() ([]domain.Club, error)
	GetClubsByName(name string) ([]*domain.Club, error)
	GetClubsByType(type_ string) ([]*domain.Club, error)
	GetClubMediaFiles(clubId int) (*domain.ClubPhoto, error)
	GetClubOrgs(clubID int) ([]domain.ClubOrg, error)
	GetClubsOrgs(clubIDs []int) ([]domain.ClubOrg, error)
	GetClubSubOrgs(clubID int) ([]domain.ClubOrg, error)
	GetAllClubOrgs() ([]domain.ClubOrg, error)
	AddClub(c *domain.Club) (int, error)
	AddOrgs(orgs []domain.ClubOrg) error
	DeleteClubWithOrgs(clubID int) error
	UpdateClub(c *domain.Club, o []domain.ClubOrg) error
}

func (s *storage) GetClub(id int) (*domain.Club, error) {
	return s.postgres.GetClub(id)
}

func (s *storage) GetAllClub() ([]domain.Club, error) {
	return s.postgres.GetAllClub()
}

func (s *storage) GetClubsByName(name string) ([]domain.Club, error) {
	return s.postgres.GetClubsByName(name)
}

func (s *storage) GetClubsByType(type_ string) ([]domain.Club, error) {
	return s.postgres.GetClubsByType(type_)
}

func (s *storage) GetClubMediaFiles(clubID int) ([]domain.ClubPhoto, error) {
	return s.postgres.GetClubMediaFiles(clubID)
}

func (s *storage) GetClubOrgs(clubID int) ([]domain.ClubOrg, error) {
	return s.postgres.GetClubOrgs(clubID)
}

func (s *storage) GetClubsOrgs(clubIDs []int) ([]domain.ClubOrg, error) {
	return s.postgres.GetClubsOrgs(clubIDs)
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

func (s *storage) DeleteClubWithOrgs(clubID int) error {
	return s.postgres.DeleteClubWithOrgs(clubID)
}

func (s *storage) UpdateClub(c *domain.Club, o []domain.ClubOrg) error {
	return s.postgres.UpdateClub(c, o)
}
