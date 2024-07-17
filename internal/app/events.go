package app

import (
	"context"
	"fmt"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type eventsServiceStorage interface {
	GetAllEvents(ctx context.Context) ([]domain.Event, error)
	GetEvent(ctx context.Context, id int) (*domain.Event, error)
	GetEventsByRange(ctx context.Context, from, to time.Time) ([]domain.Event, error)
	PostEvent(ctx context.Context, event *domain.Event) error
	DeleteEvent(ctx context.Context, id int) error
	UpdateEvent(ctx context.Context, event *domain.Event) error
	GetMediaFile(ctx context.Context, id int) (*domain.MediaFile, error)
	GetMediaFiles(ctx context.Context, ids []int) (map[int]domain.MediaFile, error)
}

type EventsService struct {
	storage eventsServiceStorage
}

func NewEventsService(storage eventsServiceStorage) *EventsService {
	return &EventsService{storage: storage}
}

func (s *EventsService) GetAllEvents(ctx context.Context) (*responses.GetAllEvents, error) {
	var res []domain.Event
	var err error

	res, err = s.storage.GetAllEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetAllEvents: %w", err)
	}

	ids := make([]int, 0, len(res))
	for _, event := range res {
		ids = append(ids, event.MediaID)
	}

	eventMediaFiles, err := s.storage.GetMediaFiles(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetEventMediaFiles: %w", err)
	}

	return mapper.MakeResponseAllEvents(res, eventMediaFiles)
}

func (s *EventsService) GetEvent(ctx context.Context, id int) (*responses.GetEvent, error) {
	res, err := s.storage.GetEvent(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetEvent: %w", err)
	}

	feedMediaFile, err := s.storage.GetMediaFile(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetEventMediaFile: %w", err)
	}

	return mapper.MakeResponseEvent(res, feedMediaFile)
}

func (s *EventsService) GetEventsByRange(ctx context.Context, from, to time.Time) (*responses.GetEventsByRange, error) {
	res, err := s.storage.GetEventsByRange(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetEventsByRange: %w", err)
	}

	ids := make([]int, 0, len(res))
	for _, event := range res {
		ids = append(ids, event.MediaID)
	}

	eventMediaFiles, err := s.storage.GetMediaFiles(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetEventMediaFiles: %w", err)
	}

	return mapper.MakeResponseEventsByRange(res, eventMediaFiles)
}

// func (s *EventsService) GetEventMemberRoles(ctx context.Context, id int) (*responses.GetEventMemberRoles, error) {

// 	return mapper.MakeResponseEventMemberRoles(res)
// }

func (s *EventsService) PostEvent(ctx context.Context, event *domain.Event) error {
	err := s.storage.PostEvent(ctx, event)
	if err != nil {
		return fmt.Errorf("can't storage.PostEvent: %w", err)
	}

	return nil
}

func (s *EventsService) DeleteEvent(ctx context.Context, id int) error {
	if err := s.storage.DeleteEvent(ctx, id); err != nil {
		return fmt.Errorf("can't storage.DeleteEvent: %w", err)
	}

	return nil
}

func (s *EventsService) UpdateEvent(ctx context.Context, event *domain.Event) error {
	if err := s.storage.UpdateEvent(ctx, event); err != nil {
		return fmt.Errorf("can't storage.UpdateEvent: %w", err)
	}

	return nil
}
