package mapper

import (
	"encoding/base64"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllFeed(f []domain.Feed) *responses.GetAllFeed {
	var feed []responses.Feed
	for _, i := range f {
		feed = append(feed,
			responses.Feed{
				ID:              i.ID,
				Title:           i.Title,
				Description:     i.Description,
				RegistrationURL: i.RegistrationURL,
				Media:           base64.Encoding{},
				CreatedBy:       i.CreatedBy,
				UpdatedAt:       i.UpdatedAt,
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
		Media:           base64.Encoding{},
		CreatedBy:       f.CreatedBy,
		UpdatedAt:       f.UpdatedAt,
	}
}

func MakeRequestPutFeed(f requests.UpdateFeed) *domain.Feed {
	return &domain.Feed{
		ID:              f.ID,
		Title:           f.Title,
		Description:     f.Description,
		RegistrationURL: f.RegistrationURL,
		CreatedBy:       f.CreatedBy,
		UpdatedAt:       f.UpdatedAt,
	}
}
