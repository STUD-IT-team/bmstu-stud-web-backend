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
}

type ClubService struct {
	storage clubStorage
}

func NewClubService(storage clubStorage) *ClubService {
	return &ClubService{storage: storage}
}

func (s *ClubService) GetClub(id int) (*responses.GetClub, error) {
	res, err := s.storage.GetClub(id)
	if err != nil {
		err = fmt.Errorf("can't storage.GetClub: %v", err)
		return nil, err
	}

	im, err := s.storage.GetMediaFile(res.LogoId)
	if err != nil {
		err = fmt.Errorf("can't storage.GetMediaFile: %v", err)
		return nil, err
	}

	return mapper.MakeResponseClub(res, im), nil
}

func (s *ClubService) GetAllClubs() (*responses.GetAllClubs, error) {
	res, err := s.storage.GetAllClub()

	if err != nil {
		err = fmt.Errorf("can't storage.GetAllClub: %v", err)
		return nil, err
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
	return mapper.MakeResponseAllClub(res, logos)
}
