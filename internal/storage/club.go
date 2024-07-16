package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type clubStorage interface {
	GetClub(ctx context.Context, id int) (*domain.Club, error)
	GetAllClub(ctx context.Context) ([]domain.Club, error)
	GetClubsByName(ctx context.Context, name string) ([]*domain.Club, error)
	GetClubsByType(ctx context.Context, type_ string) ([]*domain.Club, error)
	GetClubMediaFiles(ctx context.Context, clubId int) (*domain.ClubPhoto, error)
	GetClubOrgs(ctx context.Context, clubID int) ([]domain.ClubOrg, error)
	GetClubsOrgs(ctx context.Context, clubIDs []int) ([]domain.ClubOrg, error)
	GetClubSubOrgs(ctx context.Context, clubID int) ([]domain.ClubOrg, error)
	GetAllClubOrgs(ctx context.Context) ([]domain.ClubOrg, error)
	AddClub(ctx context.Context, c *domain.Club) (int, error)
	AddOrgs(ctx context.Context, orgs []domain.ClubOrg) error
	DeleteClubWithOrgs(ctx context.Context, clubID int) error
	UpdateClub(ctx context.Context, c *domain.Club, o []domain.ClubOrg) error
	AddClubPhotos(ctx context.Context, p []domain.ClubPhoto) error
}

func (s *storage) GetClub(ctx context.Context, id int) (*domain.Club, error) {
	return s.postgres.GetClub(ctx, id)
}

func (s *storage) GetAllClub(ctx context.Context) ([]domain.Club, error) {
	return s.postgres.GetAllClub(ctx)
}

func (s *storage) GetClubsByName(ctx context.Context, name string) ([]domain.Club, error) {
	return s.postgres.GetClubsByName(ctx, name)
}

func (s *storage) GetClubsByType(ctx context.Context, type_ string) ([]domain.Club, error) {
	return s.postgres.GetClubsByType(ctx, type_)
}

func (s *storage) GetClubMediaFiles(ctx context.Context, clubID int) ([]domain.ClubPhoto, error) {
	return s.postgres.GetClubMediaFiles(clubID)
}

func (s *storage) GetClubOrgs(ctx context.Context, clubID int) ([]domain.ClubOrg, error) {
	return s.postgres.GetClubOrgs(ctx, clubID)
}

func (s *storage) GetClubsOrgs(ctx context.Context, clubIDs []int) ([]domain.ClubOrg, error) {
	return s.postgres.GetClubsOrgs(ctx, clubIDs)
}

func (s *storage) GetClubSubOrgs(ctx context.Context, clubID int) ([]domain.ClubOrg, error) {
	return s.postgres.GetClubSubOrgs(ctx, clubID)
}

func (s *storage) GetAllClubOrgs(ctx context.Context) ([]domain.ClubOrg, error) {
	return s.postgres.GetAllClubOrgs(ctx)
}

func (s *storage) AddClub(ctx context.Context, c *domain.Club) (int, error) {
	return s.postgres.AddClub(ctx, c)
}

func (s *storage) AddOrgs(ctx context.Context, orgs []domain.ClubOrg) error {
	return s.postgres.AddOrgs(ctx, orgs)
}

func (s *storage) DeleteClubWithOrgs(ctx context.Context, clubID int) error {
	return s.postgres.DeleteClubWithOrgs(ctx, clubID)
}

func (s *storage) UpdateClub(ctx context.Context, c *domain.Club, o []domain.ClubOrg) error {
	return s.postgres.UpdateClub(ctx, c, o)
}

func (s *storage) AddClubPhotos(ctx context.Context, p []domain.ClubPhoto) error {
	return s.postgres.AddClubPhotos(ctx, p)
}
