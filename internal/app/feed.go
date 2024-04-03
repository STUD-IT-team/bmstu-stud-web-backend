package app

import (
	"context"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"

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
	GetFeedByFilterLimitAndOffset(ctx context.Context, limit, offset int) ([]domain.Feed, error)
}

type FeedService struct {
	logger  *log.Logger
	storage feedServiceStorage
}

func NewFeedService(logger *log.Logger, storage feedServiceStorage) *FeedService {
	return &FeedService{logger: logger, storage: storage}
}

func (s *FeedService) GetAllFeed(ctx context.Context, filter requests.GetFeedByFilter) (*responses.GetAllFeed, error) {
	var res []domain.Feed
	var err error

	if (filter.Limit != 0) && (filter.Offset != 0) {
		res, err = s.storage.GetFeedByFilterLimitAndOffset(ctx, filter.Limit, filter.Offset)
		if err != nil {
			log.WithError(err).Warnf("can't storage.GetFeedByFilterLimitAndOffset GetFeedByFilterLimitAndOffset")
			return nil, err
		}
	} else {
		res, err = s.storage.GetAllFeed(ctx)
		if err != nil {
			log.WithError(err).Warnf("can't storage.GetAllFeed GetAllFeed")
			return nil, err
		}
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
