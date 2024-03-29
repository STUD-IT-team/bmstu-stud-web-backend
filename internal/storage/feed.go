package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

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
