package app

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/repository"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type feedServiceSrorage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
}

type feedService struct {
	logger  *logrus.Logger
	storage feedServiceSrorage
}

func NewFeedService(logger *logrus.Logger, storage feedServiceSrorage) *feedService {
	return &feedService{logger: logger, storage: storage}
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
