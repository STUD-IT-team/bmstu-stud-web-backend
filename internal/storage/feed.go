package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type feedStorage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
	DeleteFeed(ctx context.Context, id int) error
	UpdateFeed(_ context.Context, id int, feed domain.Feed) error
	GetFeedByFilter(ctx context.Context, limit, offset int) ([]domain.Feed, error)
	GetFeedByFilterIdLastAndOffset(_ context.Context, idLast, offset int) ([]domain.Feed, error)
}

func (s *storage) GetAllFeed(ctx context.Context) ([]domain.Feed, error) {
	return s.postgres.GetAllFeed(ctx)
}

func (s *storage) GetFeed(ctx context.Context, id int) (domain.Feed, error) {
	return s.postgres.GetFeed(ctx, id)
}

func (s *storage) DeleteFeed(ctx context.Context, id int) error {
	return s.postgres.DeleteFeed(ctx, id)
}

func (s *storage) UpdateFeed(ctx context.Context, feed domain.Feed) error {
	return s.postgres.UpdateFeed(ctx, feed)
}

func (s *storage) GetFeedByFilterLimitAndOffset(ctx context.Context, limit, offset int) ([]domain.Feed, error) {
	return s.postgres.GetFeedByFilterLimitAndOffset(ctx, limit, offset)
}

func (s *storage) GetFeedByFilterIdLastAndOffset(ctx context.Context, idLast, offset int) ([]domain.Feed, error) {
	return s.postgres.GetFeedByFilterIdLastAndOffset(ctx, idLast, offset)
}
