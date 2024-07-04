package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type feedStorage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetFeedEncounters(ctx context.Context, id int) ([]domain.Encounter, error)
	GetFeedByTitle(ctx context.Context, title string) ([]domain.Feed, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
	PostFeed(ctx context.Context, feed domain.Feed) error
	DeleteFeed(ctx context.Context, id int) error
	UpdateFeed(_ context.Context, id int, feed domain.Feed) error
	// GetFeedByFilter(ctx context.Context, limit, offset int) ([]domain.Feed, error)
	// GetFeedByFilterIdLastAndOffset(_ context.Context, idLast, offset int) ([]domain.Feed, error)
}

func (s *storage) GetAllFeed(ctx context.Context) ([]domain.Feed, error) {
	return s.postgres.GetAllFeed(ctx)
}

func (s *storage) GetFeed(ctx context.Context, id int) (domain.Feed, error) {
	return s.postgres.GetFeed(ctx, id)
}

func (s *storage) GetFeedEncounters(ctx context.Context, id int) ([]domain.Encounter, error) {
	return s.postgres.GetFeedEncounters(ctx, id)
}

func (s *storage) GetFeedByTitle(ctx context.Context, title string) ([]domain.Feed, error) {
	return s.postgres.GetFeedByTitle(ctx, title)
}

func (s *storage) PostFeed(ctx context.Context, feed domain.Feed) error {
	return s.postgres.PostFeed(ctx, feed)
}

func (s *storage) DeleteFeed(ctx context.Context, id int) error {
	return s.postgres.DeleteFeed(ctx, id)
}

func (s *storage) UpdateFeed(ctx context.Context, feed domain.Feed) error {
	return s.postgres.UpdateFeed(ctx, feed)
}

// func (s *storage) GetFeedByFilterLimitAndOffset(ctx context.Context, limit, offset int) ([]domain.Feed, error) {
// 	return s.postgres.GetFeedByFilterLimitAndOffset(ctx, limit, offset)
// }

// func (s *storage) GetFeedByFilterIdLastAndOffset(ctx context.Context, idLast, offset int) ([]domain.Feed, error) {
// 	return s.postgres.GetFeedByFilterIdLastAndOffset(ctx, idLast, offset)
// }
