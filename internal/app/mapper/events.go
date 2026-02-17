package mapper

import (
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponseAllEvents(evs []domain.Event, eventMediaFiles map[int]domain.MediaFile) (*responses.GetAllEvents, error) {
	event := make([]responses.Event, 0, len(evs))
	for _, v := range evs {
		media, ok := eventMediaFiles[v.MediaID]
		if !ok {
			return nil, fmt.Errorf("can't find media for event id %v", v.MediaID)
		}
		event = append(event,
			responses.Event{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
				Prompt:      v.Prompt,
				Media:       media,
				Date:        v.Date,
				Approved:    v.Approved,
				CreatedAt:   v.CreatedAt,
				CreatedBy:   v.CreatedBy,
				RegUrl:      v.RegUrl,
				RegOpenDate: v.RegOpenDate,
				FeedbackUrl: v.FeedbackUrl,
			})
	}

	return &responses.GetAllEvents{Event: event}, nil
}

func MakeResponseEvent(v *domain.Event, eventMediaFile *domain.MediaFile) (*responses.GetEvent, error) {
	return &responses.GetEvent{
		ID:          v.ID,
		Title:       v.Title,
		Description: v.Description,
		Prompt:      v.Prompt,
		Media:       *eventMediaFile,
		Date:        v.Date,
		Approved:    v.Approved,
		CreatedAt:   v.CreatedAt,
		CreatedBy:   v.CreatedBy,
		RegUrl:      v.RegUrl,
		RegOpenDate: v.RegOpenDate,
		FeedbackUrl: v.FeedbackUrl,
	}, nil
}

func MakeResponseEventsByRange(evs []domain.Event, eventMediaFiles map[int]domain.MediaFile) (*responses.GetEventsByRange, error) {
	event := make([]responses.Event, 0, len(evs))
	for _, v := range evs {
		media, ok := eventMediaFiles[v.MediaID]
		if !ok {
			return nil, fmt.Errorf("can't find media for event id %v", v.MediaID)
		}
		event = append(event,
			responses.Event{
				ID:          v.ID,
				Title:       v.Title,
				Description: v.Description,
				Prompt:      v.Prompt,
				Media:       media,
				Date:        v.Date,
				Approved:    v.Approved,
				CreatedAt:   v.CreatedAt,
				CreatedBy:   v.CreatedBy,
				RegUrl:      v.RegUrl,
				RegOpenDate: v.RegOpenDate,
				FeedbackUrl: v.FeedbackUrl,
			})
	}

	return &responses.GetEventsByRange{Events: event}, nil
}

func MakeRequestPostEvent(v *requests.PostEvent) *domain.Event {
	return &domain.Event{
		Title:       v.Title,
		Description: v.Description,
		Prompt:      v.Prompt,
		MediaID:     v.MediaID,
		Date:        v.Date,
		Approved:    v.Approved,
		CreatedAt:   v.CreatedAt,
		CreatedBy:   v.CreatedBy,
		MainOrg:     v.CreatedBy,
		ClubID:      domain.StudSovetClubID,
		RegUrl:      v.RegUrl,
		RegOpenDate: v.RegOpenDate,
		FeedbackUrl: v.FeedbackUrl,
	}
}

func MakeRequestUpdateEvent(v *requests.UpdateEvent) *domain.Event {
	return &domain.Event{
		ID:          v.ID,
		Title:       v.Title,
		Description: v.Description,
		Prompt:      v.Prompt,
		MediaID:     v.MediaID,
		Date:        v.Date,
		Approved:    v.Approved,
		CreatedAt:   v.CreatedAt,
		CreatedBy:   v.CreatedBy,
		MainOrg:     v.CreatedBy,
		ClubID:      domain.StudSovetClubID,
		RegUrl:      v.RegUrl,
		RegOpenDate: v.RegOpenDate,
		FeedbackUrl: v.FeedbackUrl,
	}
}
