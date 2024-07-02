package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type feedServiceStorage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetAllFeedEncounters(ctx context.Context) ([]domain.Encounter, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
	GetFeedByTitle(ctx context.Context, title string) ([]domain.Feed, error)
	DeleteFeed(ctx context.Context, id int) error
	UpdateFeed(ctx context.Context, feed domain.Feed) error
	GetFeedByFilterLimitAndOffset(ctx context.Context, limit, offset int) ([]domain.Feed, error)
	GetFeedByFilterIdLastAndOffset(_ context.Context, idLast, offset int) ([]domain.Feed, error)
}

type FeedService struct {
	storage feedServiceStorage
}

func NewFeedService(storage feedServiceStorage) *FeedService {
	return &FeedService{storage: storage}
}

func (s *FeedService) GetAllFeed(ctx context.Context) (*responses.GetAllFeed, error) {
	var res []domain.Feed
	var err error

	res, err = s.storage.GetAllFeed(ctx)
	if err != nil {
		fmt.Errorf("can't storage.GetAllFeed GetAllFeed")
		return nil, err
	}

	return mapper.MakeResponseAllFeed(res), nil
}

func (s *FeedService) GetAllFeedEncounters(ctx context.Context) (*responses.GetAllFeedEncounters, error) {
	var res []domain.Encounter
	var err error

	res, err = s.storage.GetAllFeedEncounters(ctx)
	if err != nil {
		fmt.Errorf("can't storage.GetAllFeedEncounters GetAllFeedEncounters")
		return nil, err
	}

	return mapper.MakeResponseAllFeedEncounters(res), nil
}

func (s *FeedService) GetFeedByTitle(
	ctx context.Context,
	filter requests.GetFeedByTitle,
) (*responses.GetAllFeedByTitle, error) {
	var res []domain.Feed
	var err error

	res, err = s.storage.GetFeedByTitle(ctx, filter.Search)
	if err != nil {
		fmt.Errorf("can't storage.GetFeedByTitle GetFeedByTitle")
		return nil, err
	}

	return mapper.MakeResponseFeedByTitle(res), nil
}

func (s *FeedService) GetFeedByFilter(
	ctx context.Context,
	filter requests.GetFeedByFilter,
) (*responses.GetAllFeed, error) {
	var res []domain.Feed
	var err error

	if filter.Limit.IsPresent() && filter.Offset.IsPresent() {
		res, err = s.storage.GetFeedByFilterLimitAndOffset(ctx, filter.Limit.MustGet(), filter.Offset.MustGet())
		if err != nil {
			fmt.Errorf("can't storage.GetFeedByFilterLimitAndOffset GetFeedByFilterLimitAndOffset")
			return nil, err
		}
	} else if filter.IdLast.IsPresent() && filter.Offset.IsPresent() {
		res, err = s.storage.GetFeedByFilterIdLastAndOffset(ctx, filter.IdLast.MustGet(), filter.Offset.MustGet())
		if err != nil {
			return nil, fmt.Errorf("can't storage.GetFeedByFilterIdLastAndOffset GetFeedByFilterIdLastAndOffset")
		}
	} else {
		res, err = s.storage.GetAllFeed(ctx)
		if err != nil {
			return nil, fmt.Errorf("can't storage.GetAllFeed GetAllFeed")
		}
	}

	return mapper.MakeResponseAllFeed(res), nil
}

func (s *FeedService) GetFeed(ctx context.Context, id int) (*responses.GetFeed, error) {
	res, err := s.storage.GetFeed(ctx, id)
	if err != nil {
		fmt.Errorf("can't storage.GetFeed GetFeed")
		return nil, err
	}

	return mapper.MakeResponseFeed(res), nil
}

func (s *FeedService) DeleteFeed(ctx context.Context, id int) error {
	if err := s.storage.DeleteFeed(ctx, id); err != nil {
		fmt.Errorf("can't storage.DeleteFeed DeleteFeed")
		return err
	}

	return nil
}

func (s *FeedService) UpdateFeed(ctx context.Context, feed domain.Feed) error {
	if err := s.storage.UpdateFeed(ctx, feed); err != nil {
		fmt.Errorf("can't storage.UpdateFeed UpdateFeed")
		return err
	}

	return nil
}
