package app

import (
	"context"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"

	log "github.com/sirupsen/logrus"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type eventServiceStorage interface {
	CreateNewEvent(ctx context.Context, event domain.Event) error
	GetAllEventsByPeriod(ctx context.Context, startDate time.Time, endDate time.Time) ([]domain.Event, error)
	GetEventByID(ctx context.Context, id int) (domain.Event, error)
	UpdateByID(ctx context.Context, event domain.Event) error
}

type EventService struct {
	logger  *log.Logger
	storage eventServiceStorage
}

func NewEventService(logger *log.Logger, storage eventServiceStorage) *EventService {
	return &EventService{
		logger:  logger,
		storage: storage,
	}
}
func (s *EventService) GetAllEvents(ctx context.Context, request.GetAllEvents) ([]domain.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventByID(ctx context.Context, ID int) (responses.GetEventByID, error) {
	return responses.GetEventByID{}, nil
}

func (s *EventService) UpdateEventByID(ctx context.Context, event domain.Event) error {
	return nil
}
func (s *EventService) CreateEvent(ctx context.Context, event domain.Event) error {
	return nil
}
