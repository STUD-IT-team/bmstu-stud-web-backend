package app

import (
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type clubStorage interface {
	GetClub(id int) (*domain.Club, error)
	GetAllClub() ([]domain.Club, error)
	GetMediaFile(id int) (*domain.MediaFile, error)
	GetMediaFiles(ids []int) (map[int]domain.MediaFile, error)
	GetClubOrgs(clubID int) ([]domain.ClubOrg, error)
	GetClubSubOrgs(clubID int) ([]domain.ClubOrg, error)
	GetAllClubOrgs() ([]domain.ClubOrg, error)
}

type ClubService struct {
	storage clubStorage
}

func NewClubService(storage clubStorage) *ClubService {
	return &ClubService{storage: storage}
}

func (s *ClubService) GetClub(id int) (*responses.GetClub, error) {
	club, err := s.storage.GetClub(id)
	if err != nil {
		err = fmt.Errorf("can't storage.GetClub: %v", err)
		return nil, err
	}

	mainOrgs, err := s.storage.GetClubOrgs(id)
	if err != nil {
		err = fmt.Errorf("can't storage.GetClubOrgs: %v", err)
		return nil, err
	}

	subOrgs, err := s.storage.GetClubSubOrgs(id)
	if err != nil {
		err = fmt.Errorf("can't storage.GetClubSubOrgs: %v", err)
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
		err = fmt.Errorf("can't storage.GetMediaFiles: %v", err)
		return nil, err
	}

	return mapper.MakeResponseClub(club, &mainOrgs, &subOrgs, &ims)
}

func (s *ClubService) GetAllClubs() (*responses.GetAllClubs, error) {
	res, err := s.storage.GetAllClub()

	if err != nil {
		err = fmt.Errorf("can't storage.GetAllClub: %v", err)
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
		err = fmt.Errorf("can't storage.GetMediaFiles: %v", err)
		return nil, err
	}

	orgs, err := s.storage.GetAllClubOrgs()
	if err != nil {
		err = fmt.Errorf("can't storage.GetAllClubOrgs: %v", err)
		return nil, err
	}
	return mapper.MakeResponseAllClub(res, logos, orgs)
}
