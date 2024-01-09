package app

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"

	log "github.com/sirupsen/logrus"
)

type feedServiceStorage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
}

type FeedService struct {
	logger  *log.Logger
	storage feedServiceStorage
}

func NewFeedService(logger *log.Logger, storage feedServiceStorage) *FeedService {
	return &FeedService{logger: logger, storage: storage}
}

func (s *FeedService) GetAllFeed(ctx context.Context) (*responses.GetAllFeed, error) {
	res, err := s.storage.GetAllFeed(ctx)
	if err != nil {
		log.WithError(err).Warnf("can't storage.GetAllFeed GetAllFeed")
		return nil, err
	}
	return mapper.MakeResponseAllFeed(res), nil
}

func (s *FeedService) GetFeed(ctx context.Context, id int) (*responses.GetFeed, error) {
	res, err := s.storage.GetFeed(ctx, id)
	if err != nil {
		log.WithError(err).Warnf("can't storage.GetFeed GetFeed")
		return nil, err
	}

	return mapper.MakeResponseFeed(res), nil
}
