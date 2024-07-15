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
	clubsService   *ClubService
	membersService *MembersService
	faqService     *FAQService
}

func NewAPI(logger *logrus.Logger, feed *FeedService, events *EventsService, membs *MembersService, club *ClubService, guard *GuardService, faq *FAQService) API {
	return &apiService{
		logger:         logger,
		feedService:    feed,
		guardService:   guard,
		eventsService:  events,
		clubsService:   club,
		membersService: membs,
		faqService:     faq,
	}
}

func (s *apiService) Echo() *responses.GetEcho {
	return mapper.CreateEcho()
}
