package app

import (
	"github.com/sirupsen/logrus"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type API interface {
	Echo() *responses.GetEcho
}

type apiService struct {
	logger         *logrus.Logger
	feedService    *FeedService
	guardService   *GuardService
	eventsService  *EventsService
	membersService *MembersService
}

func NewAPI(logger *logrus.Logger, feed *FeedService, events *EventsService, membs *MembersService, guard *GuardService) API {
	return &apiService{
		logger:         logger,
		feedService:    feed,
		guardService:   guard,
		eventsService:  events,
		membersService: membs,
	}
}

func (s *apiService) Echo() *responses.GetEcho {
	return mapper.CreateEcho()
}
