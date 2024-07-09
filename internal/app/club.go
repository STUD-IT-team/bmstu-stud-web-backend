package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type clubStorage interface {
	GetClub(ctx context.Context, id int) (*domain.Club, error)
	GetAllClub(ctx context.Context) ([]domain.Club, error)
	GetClubsByName(ctx context.Context, name string) ([]domain.Club, error)
	GetClubsByType(ctx context.Context, type_ string) ([]domain.Club, error)
	GetMediaFile(id int) (*domain.MediaFile, error)
	GetMediaFiles(ids []int) (map[int]domain.MediaFile, error)
	GetClubMediaFiles(ctx context.Context, clubID int) ([]domain.ClubPhoto, error)
	GetClubOrgs(ctx context.Context, clubID int) ([]domain.ClubOrg, error)
	GetClubsOrgs(ctx context.Context, clubIDs []int) ([]domain.ClubOrg, error)
	GetClubSubOrgs(ctx context.Context, clubID int) ([]domain.ClubOrg, error)
	GetAllClubOrgs(ctx context.Context) ([]domain.ClubOrg, error)
	AddClub(ctx context.Context, c *domain.Club) (int, error)
	AddOrgs(ctx context.Context, orgs []domain.ClubOrg) error
	DeleteClubWithOrgs(ctx context.Context, clubID int) error
	UpdateClub(ctx context.Context, c *domain.Club, o []domain.ClubOrg) error
}

type ClubService struct {
	storage clubStorage
}

func NewClubService(storage clubStorage) *ClubService {
	return &ClubService{storage: storage}
}

func (s *ClubService) GetClub(ctx context.Context, id int) (*responses.GetClub, error) {
	club, err := s.storage.GetClub(ctx, id)
	if err != nil {
		err = fmt.Errorf("can't storage.GetClub: %w", err)
		return nil, err
	}

	mainOrgs, err := s.storage.GetClubOrgs(ctx, id)
	if err != nil {
		err = fmt.Errorf("can't storage.GetClubOrgs: %w", err)
		return nil, err
	}

	subOrgs, err := s.storage.GetClubSubOrgs(ctx, id)
	if err != nil {
		err = fmt.Errorf("can't storage.GetClubSubOrgs: %w", err)
		return nil, err
	}

	ids := make([]int, 0, len(mainOrgs)+len(subOrgs)+1)
	for _, org := range mainOrgs {
		ids = append(ids, org.MediaID)
	}
	for _, org := range subOrgs {
		ids = append(ids, org.MediaID)
	}
	ids = append(ids, club.LogoId)

	ims, err := s.storage.GetMediaFiles(ids)
	if err != nil {
		err = fmt.Errorf("can't storage.GetMediaFiles: %w", err)
		return nil, err
	}

	return mapper.MakeResponseClub(club, &mainOrgs, &subOrgs, &ims)
}

func (s *ClubService) GetClubsByName(ctx context.Context, name string) (*responses.GetClubsByName, error) {
	res, err := s.storage.GetClubsByName(ctx, name)
	if err != nil {
		err = fmt.Errorf("can't storage.GetClubsByName: %w", err)
		return nil, err
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("no club found")
	}

	ids := make([]int, 0, len(res))
	for _, club := range res {
		ids = append(ids, club.LogoId)
	}

	logos, err := s.storage.GetMediaFiles(ids)
	if err != nil {
		err = fmt.Errorf("can't storage.GetMediaFiles: %w", err)
		return nil, err
	}

	orgs, err := s.storage.GetAllClubOrgs(ctx)
	if err != nil {
		err = fmt.Errorf("can't storage.GetAllClubOrgs: %w", err)
		return nil, err
	}
	resp, err := mapper.MakeResponseAllClub(res, logos, orgs)
	clubs := resp.Clubs
	return &responses.GetClubsByName{Clubs: clubs}, err
}

func (s *ClubService) GetClubsByType(ctx context.Context, type_ string) (*responses.GetClubsByType, error) {
	res, err := s.storage.GetClubsByType(ctx, type_)
	if err != nil {
		err = fmt.Errorf("can't storage.GetClubsByName: %w", err)
		return nil, err
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("no club found")
	}

	ids := make([]int, 0, len(res))
	for _, club := range res {
		ids = append(ids, club.LogoId)
	}

	logos, err := s.storage.GetMediaFiles(ids)
	if err != nil {
		err = fmt.Errorf("can't storage.GetMediaFiles: %w", err)
		return nil, err
	}

	orgs, err := s.storage.GetAllClubOrgs(ctx)
	if err != nil {
		err = fmt.Errorf("can't storage.GetAllClubOrgs: %w", err)
		return nil, err
	}
	resp, err := mapper.MakeResponseAllClub(res, logos, orgs)
	clubs := resp.Clubs
	return &responses.GetClubsByType{Clubs: clubs}, err
}

func (s *ClubService) GetAllClubs(ctx context.Context) (*responses.GetAllClubs, error) {
	res, err := s.storage.GetAllClub(ctx)

	if err != nil {
		err = fmt.Errorf("can't storage.GetAllClub: %w", err)
		return nil, err
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("no club found")
	}

	ids := make([]int, 0, len(res))
	for _, club := range res {
		ids = append(ids, club.LogoId)
	}

	logos, err := s.storage.GetMediaFiles(ids)
	if err != nil {
		err = fmt.Errorf("can't storage.GetMediaFiles: %w", err)
		return nil, err
	}

	orgs, err := s.storage.GetAllClubOrgs(ctx)
	if err != nil {
		err = fmt.Errorf("can't storage.GetAllClubOrgs: %w", err)
		return nil, err
	}
	return mapper.MakeResponseAllClub(res, logos, orgs)
}

func (s *ClubService) GetClubMembers(ctx context.Context, clubID int) (*responses.GetClubMembers, error) {
	orgs, err := s.storage.GetClubOrgs(ctx, clubID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetClubOrgs: %w", err)
	}

	subOrgs, err := s.storage.GetClubSubOrgs(ctx, clubID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetClubSubOrgs: %w", err)
	}

	if len(orgs)+len(subOrgs) == 0 {
		return nil, fmt.Errorf("no club members found")
	}

	ids := make([]int, 0, len(orgs)+len(subOrgs))
	for _, org := range orgs {
		ids = append(ids, org.MediaID)
	}
	for _, org := range subOrgs {
		ids = append(ids, org.MediaID)
	}

	media, err := s.storage.GetMediaFiles(ids)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetMediaFiles: %w", err)
	}
	return mapper.MakeResponseClubMembers(clubID, orgs, subOrgs, media)
}

func (s *ClubService) PostClub(ctx context.Context, req *requests.PostClub) error {
	club, orgs, err := mapper.ParsePostClub(req)
	if err != nil {
		return fmt.Errorf("can't mapper.PostClub: %w", err)
	}
	clubID, err := s.storage.AddClub(ctx, club)
	if err != nil {
		return fmt.Errorf("can't storage.AddClub: %w", err)
	}

	for i := range orgs {
		orgs[i].ClubID = clubID
	}

	err = s.storage.AddOrgs(ctx, orgs)
	if err != nil {
		return fmt.Errorf("can't storage.AddOrgs: %w", err)
	}

	return nil
}

func (s *ClubService) DeleteClub(ctx context.Context, clubID int) error {
	err := s.storage.DeleteClubWithOrgs(ctx, clubID)
	if err != nil {
		return fmt.Errorf("can't storage.DeleteClubWithOrgs: %w", err)
	}

	return nil
}

func (s *ClubService) UpdateClub(ctx context.Context, req *requests.UpdateClub) error {
	club, orgs, err := mapper.ParseUpdateClub(req)
	if err != nil {
		return fmt.Errorf("can't mapper.PostClub: %w", err)
	}
	err = s.storage.UpdateClub(ctx, club, orgs)
	if err != nil {
		return fmt.Errorf("can't storage.UpdateClub: %w", err)
	}

	return nil
}

func (s *ClubService) GetClubMediaFiles(ctx context.Context, clubID int) (*responses.GetClubMedia, error) {
	res, err := s.storage.GetClubMediaFiles(ctx, clubID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetClubMediaFiles: %w", err)
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("no club photo found")
	}

	return mapper.MakeResponseClubMediaFiles(clubID, res)
}
