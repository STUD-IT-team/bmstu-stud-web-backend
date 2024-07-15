package storage

import (
	"context"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type eventStorage interface {
	GetAllEvents(ctx context.Context) ([]domain.Event, error)
	GetEvent(ctx context.Context, id int) (*domain.Event, error)
	GetEventsByRange(ctx context.Context, from, to time.Time) ([]domain.Event, error)
	PostEvent(ctx context.Context, event *domain.Event) error
	DeleteEvent(ctx context.Context, id int) error
	UpdateEvent(cx context.Context, event *domain.Event) error
}

func (s *storage) GetAllEvents(ctx context.Context) ([]domain.Event, error) {
	return s.postgres.GetAllEvents(ctx)
}

func (s *storage) GetEvent(ctx context.Context, id int) (*domain.Event, error) {
	return s.postgres.GetEvent(ctx, id)
}

func (s *storage) GetEventsByRange(ctx context.Context, from, to time.Time) ([]domain.Event, error) {
	return s.postgres.GetEventsByRange(ctx, from, to)
}

func (s *storage) PostEvent(ctx context.Context, event *domain.Event) error {
	return s.postgres.PostEvent(ctx, event)
}

func (s *storage) DeleteEvent(ctx context.Context, id int) error {
	return s.postgres.DeleteEvent(ctx, id)
}

func (s *storage) UpdateEvent(ctx context.Context, event *domain.Event) error {
	return s.postgres.UpdateEvent(ctx, event)
}
