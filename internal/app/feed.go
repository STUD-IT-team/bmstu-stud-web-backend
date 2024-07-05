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
	GetFeedEncounters(ctx context.Context, id int) ([]domain.Encounter, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
	GetFeedByTitle(ctx context.Context, title string) ([]domain.Feed, error)
	PostFeed(ctx context.Context, feed domain.Feed) error
	DeleteFeed(ctx context.Context, id int) error
	UpdateFeed(ctx context.Context, feed domain.Feed) error
	GetMediaFile(id int) (*domain.MediaFile, error)
	GetMediaFiles(ids []int) (map[int]domain.MediaFile, error)
	// GetFeedByFilterLimitAndOffset(ctx context.Context, limit, offset int) ([]domain.Feed, error)
	// GetFeedByFilterIdLastAndOffset(_ context.Context, idLast, offset int) ([]domain.Feed, error)
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
		return nil, fmt.Errorf("can't storage.GetAllFeed: %v", err)
	}

	ids := make([]int, 0, len(res))
	for _, feed := range res {
		ids = append(ids, feed.MediaID)
	}

	feedMediaFiles, err := s.storage.GetMediaFiles(ids)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetFeedMediaFiles: %v", err)
	}

	return mapper.MakeResponseAllFeed(res, feedMediaFiles)
}

func (s *FeedService) GetFeed(ctx context.Context, id int) (*responses.GetFeed, error) {
	res, err := s.storage.GetFeed(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetFeed: %v", err)
	}

	feedMediaFile, err := s.storage.GetMediaFile(id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetFeedMediaFile: %v", err)
	}

	return mapper.MakeResponseFeed(&res, feedMediaFile)
}

func (s *FeedService) GetFeedEncounters(ctx context.Context, id int) (*responses.GetFeedEncounters, error) {
	var res []domain.Encounter
	var err error

	res, err = s.storage.GetFeedEncounters(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetFeedEncounters: %v", err)
	}

	return mapper.MakeResponseFeedEncounters(res)
}

func (s *FeedService) GetFeedByTitle(
	ctx context.Context,
	filter requests.GetFeedByTitle,
) (*responses.GetFeedByTitle, error) {
	var res []domain.Feed
	var err error

	res, err = s.storage.GetFeedByTitle(ctx, filter.Search)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetFeedByTitle: %v", err)
	}

	ids := make([]int, 0, len(res))
	for _, feed := range res {
		ids = append(ids, feed.MediaID)
	}

	feedMediaFiles, err := s.storage.GetMediaFiles(ids)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetFeedMediaFiles: %v", err)
	}

	return mapper.MakeResponseFeedByTitle(res, feedMediaFiles)
}

func (s *FeedService) PostFeed(ctx context.Context, feed domain.Feed) error {
	err := s.storage.PostFeed(ctx, feed)
	if err != nil {
		return fmt.Errorf("can't storage.PostFeed: %v", err)
	}

	return nil
}

func (s *FeedService) DeleteFeed(ctx context.Context, id int) error {
	if err := s.storage.DeleteFeed(ctx, id); err != nil {
		return fmt.Errorf("can't storage.DeleteFeed: %v", err)
	}

	return nil
}

func (s *FeedService) UpdateFeed(ctx context.Context, feed domain.Feed) error {
	if err := s.storage.UpdateFeed(ctx, feed); err != nil {
		return fmt.Errorf("can't storage.UpdateFeed: %v", err)
	}

	return nil
}

// func (s *FeedService) GetFeedByFilter(
// 	ctx context.Context,
// 	filter requests.GetFeedByFilter,
// ) (*responses.GetAllFeed, error) {
// 	var res []domain.Feed
// 	var err error

// 	if filter.Limit.IsPresent() && filter.Offset.IsPresent() {
// 		res, err = s.storage.GetFeedByFilterLimitAndOffset(ctx, filter.Limit.MustGet(), filter.Offset.MustGet())
// 		if err != nil {
// 			return nil, fmt.Errorf("can't storage.GetFeedByFilterLimitAndOffset: %v", err)
// 		}
// 	} else if filter.IdLast.IsPresent() && filter.Offset.IsPresent() {
// 		res, err = s.storage.GetFeedByFilterIdLastAndOffset(ctx, filter.IdLast.MustGet(), filter.Offset.MustGet())
// 		if err != nil {
// 			return nil, fmt.Errorf("can't storage.GetFeedByFilterIdLastAndOffset: %v", err)
// 		}
// 	} else {
// 		res, err = s.storage.GetAllFeed(ctx)
// 		if err != nil {
// 			return nil, fmt.Errorf("can't storage.GetAllFeed: %v", err)
// 		}
// 	}

// 	return mapper.MakeResponseAllFeed(res), nil
// }
