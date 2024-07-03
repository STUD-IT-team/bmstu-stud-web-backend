package mapper

import (
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
				Approved:    v.Approved,
				MediaID:     v.MediaID,
				VkPostUrl:   v.VkPostUrl,
				UpdatedAt:   v.UpdatedAt,
				CreatedAt:   v.CreatedAt,
				CreatedBy:   v.CreatedBy,
				Views:       v.Views,
			})
	}

	return &responses.GetAllFeed{Feed: feed}
}

func MakeResponseFeedEncounters(f []domain.Encounter) *responses.GetFeedEncounters {
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

	return &responses.GetFeedEncounters{Encounters: encounters}
}

func MakeResponseFeedByTitle(f []domain.Feed) *responses.GetAllFeedByTitle {
	feed := make([]responses.Feed, 0, len(f))
	for _, v := range f {
		feed = append(feed,
			responses.Feed{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
				Approved:    v.Approved,
				MediaID:     v.MediaID,
				VkPostUrl:   v.VkPostUrl,
				UpdatedAt:   v.UpdatedAt,
				CreatedAt:   v.CreatedAt,
				CreatedBy:   v.CreatedBy,
				Views:       v.Views,
			})
	}

	return &responses.GetAllFeedByTitle{Feed: feed}
}

func MakeResponseFeed(f domain.Feed) *responses.GetFeed {
	return &responses.GetFeed{
		ID:          f.ID,
		Title:       f.Title,
		Description: f.Description,
		Approved:    f.Approved,
		MediaID:     f.MediaID,
		VkPostUrl:   f.VkPostUrl,
		UpdatedAt:   f.UpdatedAt,
		CreatedAt:   f.CreatedAt,
		CreatedBy:   f.CreatedBy,
		Views:       f.Views,
	}
}

func MakeRequestPostFeed(f requests.PostFeed) *domain.Feed {
	return &domain.Feed{
		Title:       f.Title,
		Description: f.Description,
		Approved:    f.Approved,
		MediaID:     f.MediaID,
		VkPostUrl:   f.VkPostUrl,
		UpdatedAt:   f.UpdatedAt,
		CreatedAt:   f.CreatedAt,
		CreatedBy:   f.CreatedBy,
		Views:       f.Views,
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
