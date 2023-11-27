package app

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type FeedService struct {
	repository repository.IFeedRepository
}

func NewFeedService(repository repository.IFeedRepository) *FeedService {
	return &FeedService{repository: repository}
}

func (s *FeedService) GetAllFeed() ([]responses.Feed, error) {
	return s.repository.GetAllFeed()
}

func (s *FeedService) GetFeed(id int) (responses.Feed, error) {
	return s.repository.GetFeed(id)
}
