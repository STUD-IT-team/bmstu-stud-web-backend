package app

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type feedServiceStorage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
	DeleteFeed(ctx context.Context, id int) error
	UpdateFeed(ctx context.Context, feed domain.Feed) error
	GetFeedByFilter(ctx context.Context, limit, offset int) ([]domain.Feed, error)
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

func (s *FeedService) DeleteFeed(ctx context.Context, id int) error {
	if err := s.storage.DeleteFeed(ctx, id); err != nil {
		log.WithError(err).Warnf("can't storage.DeleteFeed DeleteFeed")
		return err
	}

	return nil
}

func (s *FeedService) UpdateFeed(ctx context.Context, feed domain.Feed) error {
	if err := s.storage.UpdateFeed(ctx, feed); err != nil {
		log.WithError(err).Warnf("can't storage.UpdateFeed UpdateFeed")
		return err
	}

	return nil
}

func (s *FeedService) GetFeedByFilter(ctx context.Context, limit, offset int) (*responses.GetAllFeed, error) {
	res, err := s.storage.GetFeedByFilter(ctx, limit, offset)
	if err != nil {
		log.WithError(err).Warnf("can't storage.GetFeedByFilter GetFeedByFilter")
		return nil, err
	}

	return mapper.MakeResponseAllFeed(res), nil
}
