package mapper

import (
	"encoding/base64"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllFeed(f []domain.Feed) *responses.GetAllFeed {
	feed := make([]responses.Feed, 0, len(f))
	for _, v := range f {
		feed = append(feed,
			responses.Feed{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
				Media:       base64.Encoding{},
				CreatedBy:   v.CreatedBy,
				UpdatedAt:   v.UpdatedAt,
			})
	}

	return &responses.GetAllFeed{Feed: feed}
}

func MakeResponseAllFeedEncounters(f []domain.Encounter) *responses.GetAllFeedEncounters {
	encounters := make([]responses.Encounter, 0, len(f))
	for _, v := range f {
		encounters = append(encounters,
			responses.Encounter{
				ID:          v.ID,
				Count:       v.Count,
				Description: v.Description,
				ClubID:      v.ClubID,
			})
	}

	return &responses.GetAllFeedEncounters{Encounters: encounters}
}

func MakeResponseFeedByTitle(f []domain.Feed) *responses.GetAllFeedByTitle {
	feed := make([]responses.Feed, 0, len(f))
	for _, v := range f {
		feed = append(feed,
			responses.Feed{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
			})
	}

	return &responses.GetAllFeedByTitle{Feed: feed}
}

func MakeResponseFeed(f domain.Feed) *responses.GetFeed {
	return &responses.GetFeed{
		ID:          f.ID,
		Title:       f.Title,
		Description: f.Description,
		Media:       base64.Encoding{},
		CreatedBy:   f.CreatedBy,
		UpdatedAt:   f.UpdatedAt,
	}
}

func MakeRequestPutFeed(f requests.UpdateFeed) *domain.Feed {
	return &domain.Feed{
		ID:          f.ID,
		Title:       f.Title,
		Description: f.Description,
		CreatedBy:   f.CreatedBy,
		UpdatedAt:   f.UpdatedAt,
	}
}
