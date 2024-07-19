package mapper

import (
	"fmt"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllFeed(f []domain.Feed, feedMediaFiles map[int]domain.MediaFile) (*responses.GetAllFeed, error) {
	feed := make([]responses.Feed, 0, len(f))
	for _, v := range f {
		media, ok := feedMediaFiles[v.MediaID]
		if !ok {
			return nil, fmt.Errorf("can't find media for feed id %v", v.MediaID)
		}
		feed = append(feed,
			responses.Feed{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
				Approved:    v.Approved,
				Media:       media,
				VkPostUrl:   v.VkPostUrl,
				UpdatedAt:   v.UpdatedAt,
				CreatedAt:   v.CreatedAt,
				CreatedBy:   v.CreatedBy,
				Views:       v.Views,
			})
	}

	return &responses.GetAllFeed{Feed: feed}, nil
}

func MakeResponseFeed(f *domain.Feed, feedMediaFile *domain.MediaFile) (*responses.GetFeed, error) {
	return &responses.GetFeed{
		ID:          f.ID,
		Title:       f.Title,
		Description: f.Description,
		Approved:    f.Approved,
		Media:       *feedMediaFile,
		VkPostUrl:   f.VkPostUrl,
		UpdatedAt:   f.UpdatedAt,
		CreatedAt:   f.CreatedAt,
		CreatedBy:   f.CreatedBy,
		Views:       f.Views,
	}, nil
}

func MakeResponseFeedEncounters(f []domain.Encounter) (*responses.GetFeedEncounters, error) {
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

	return &responses.GetFeedEncounters{Encounters: encounters}, nil
}

func MakeResponseFeedByTitle(f []domain.Feed, feedMediaFiles map[int]domain.MediaFile) (*responses.GetFeedByTitle, error) {
	feed := make([]responses.Feed, 0, len(f))
	for _, v := range f {
		media, ok := feedMediaFiles[v.MediaID]
		if !ok {
			return nil, fmt.Errorf("can't find media for feed id %v", v.MediaID)
		}
		feed = append(feed,
			responses.Feed{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
				Approved:    v.Approved,
				Media:       media,
				VkPostUrl:   v.VkPostUrl,
				UpdatedAt:   v.UpdatedAt,
				CreatedAt:   v.CreatedAt,
				CreatedBy:   v.CreatedBy,
				Views:       v.Views,
			})
	}

	return &responses.GetFeedByTitle{Feed: feed}, nil
}

func MakeRequestPostFeed(f *requests.PostFeed, currTime time.Time, id int) *domain.Feed {
	return &domain.Feed{
		Title:       f.Title,
		Description: f.Description,
		Approved:    f.Approved,
		MediaID:     f.MediaID,
		VkPostUrl:   f.VkPostUrl,
		CreatedAt:   currTime,
		CreatedBy:   id,
		UpdatedAt:   currTime,
	}
}

func MakeRequestUpdateFeed(f *requests.UpdateFeed, currTime time.Time, id int) *domain.Feed {
	return &domain.Feed{
		ID:          f.ID,
		Title:       f.Title,
		Description: f.Description,
		Approved:    f.Approved,
		MediaID:     f.MediaID,
		VkPostUrl:   f.VkPostUrl,
		CreatedAt:   currTime,
		CreatedBy:   id,
		UpdatedAt:   currTime,
		Views:       f.Views,
	}
}

func MakeRequestPostEncounter(f *requests.PostEncounter) *domain.Encounter {
	return &domain.Encounter{
		Count:       f.Count,
		Description: f.Description,
		ClubID:      f.ClubID,
	}
}

func MakeRequestUpdateEncounter(f *requests.UpdateEncounter) *domain.Encounter {
	return &domain.Encounter{
		ID:          f.ID,
		Count:       f.Count,
		Description: f.Description,
		ClubID:      f.ClubID,
	}
}
