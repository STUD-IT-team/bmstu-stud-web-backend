package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllFeed(dom []domain.Feed) *responses.GetAllFeed {
	feed := []responses.Feed{}
	for _, i := range dom {
		feed = append(feed, responses.Feed{
			ID: i.ID, Title: i.Title, Description: i.Description,
		})
	}
	return &responses.GetAllFeed{Feed: feed}
}

func MakeResponseFeed(dom domain.Feed) *responses.GetFeed {
	return &responses.GetFeed{
		ID: dom.ID, Title: dom.Title, Description: dom.Description, RegistationURL: dom.RegistationURL,
	}
}
