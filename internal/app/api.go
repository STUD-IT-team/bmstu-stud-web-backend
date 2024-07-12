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
	logger                   *logrus.Logger
	feedService              *FeedService
	guardService             *GuardService
	eventsService            *EventsService
	clubsService             *ClubService
	membersService           *MembersService
	documentsService         *DocumentsService
	documentCategoriesSevice *DocumentCategoriesService
}

func NewAPI(logger *logrus.Logger, feed *FeedService, events *EventsService,
	membs *MembersService, club *ClubService, guard *GuardService,
	docs *DocumentsService, cats *DocumentCategoriesService) API {
	return &apiService{
		logger:                   logger,
		feedService:              feed,
		guardService:             guard,
		eventsService:            events,
		clubsService:             club,
		membersService:           membs,
		documentsService:         docs,
		documentCategoriesSevice: cats,
	}
}

func (s *apiService) Echo() *responses.GetEcho {
	return mapper.CreateEcho()
}
