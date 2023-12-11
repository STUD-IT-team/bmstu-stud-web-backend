package app

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/storage"
)

type FeedServiceSrorage interface {
	GetAllFeed() ([]responses.GetAllFeed, error)
	GetFeed(id int) (responses.GetFeed, error)
}

type feedService struct {
	storage storage.Storage
}

func NewFeedService(storage storage.Storage) *feedService {
	return &feedService{storage: storage}
}

func (s *feedService) GetAllFeed() ([]responses.GetAllFeed, error) {
	res, err := s.storage.GetAllFeed()
	if err != nil {
		return nil, err
	}
	return *mapper.MakeResponseAllFeed(res), nil
}

func (s *feedService) GetFeed(id int) (responses.GetFeed, error) {
	res, err := s.storage.GetFeed(id)
	if err != nil {
		return responses.GetFeed{}, err
	}
	return *mapper.MakeResponseFeed(res), nil
}
