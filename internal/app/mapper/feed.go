package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllFeed(f []domain.Feed) *responses.GetAllFeed {
	feed := []responses.Feed{}
	for _, i := range f {
		feed = append(feed,
			responses.Feed{
				ID:          i.ID,
				Title:       i.Title,
				Description: i.Description,
			})
	}

	return &responses.GetAllFeed{Feed: feed}
}

func MakeResponseFeed(f domain.Feed) *responses.GetFeed {
	return &responses.GetFeed{
		ID:              f.ID,
		Title:           f.Title,
		Description:     f.Description,
		RegistrationURL: f.RegistrationURL,
	}
}
