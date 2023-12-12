package app

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/storage"
	log "github.com/sirupsen/logrus"
)

type FeedServiceSrorage interface {
	GetAllFeed(ctx context.Context) ([]responses.GetAllFeed, error)
	GetFeed(ctx context.Context, id int) (responses.GetFeed, error)
}

type feedService struct {
	storage storage.Storage
}

func NewFeedService(repository repository.IFeedRepository) *FeedService {
	return &FeedService{repository: repository}
}

func (s *feedService) GetAllFeed(ctx context.Context) ([]responses.GetAllFeed, error) {
	res, err := s.storage.GetAllFeed(ctx)

	if err != nil {
		log.WithField("", "GetAllFeed").Error(err)
		return nil, err
	}
	return *mapper.MakeResponseAllFeed(res), nil
}

func (s *feedService) GetFeed(ctx context.Context, id int) (responses.GetFeed, error) {
	res, err := s.storage.GetFeed(ctx, id)
	if err != nil {
		log.WithField("", "GetFeed").Error(err)
		return responses.GetFeed{}, err
	}
	return *mapper.MakeResponseFeed(res), nil
}
