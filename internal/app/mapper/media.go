package mapper

import (
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func MakeResponsePostMedia(id int) *responses.PostMedia {
	return &responses.PostMedia{
		ID: id,
	}
}

func MakeResponseGetDefaultMedia(defaultMedia *domain.DefaultMedia, media *domain.MediaFile) (*responses.GetDefaultMedia, error) {
	if defaultMedia == nil || media == nil {
		return nil, fmt.Errorf("got nil defaultMedia or media")
	}
	if defaultMedia.MediaID != media.ID {
		return nil, fmt.Errorf("got not matching media_id in defaultMedia and media id: %v != %v", defaultMedia.MediaID, media.ID)
	}

	resp := responses.GetDefaultMedia{}
	resp.DefaultID = defaultMedia.ID
	resp.ID = media.ID
	resp.Key = media.Key
	resp.Name = media.Name

	return &resp, nil
}

func MakeResponseAllDefaultMedia(defaultMedias []domain.DefaultMedia, mediaFiles map[int]domain.MediaFile) (*responses.GetAllDefaultMedia, error) {
	resp := responses.GetAllDefaultMedia{}
	for _, defaultMedia := range defaultMedias {
		media, ok := mediaFiles[defaultMedia.MediaID]
		if !ok {
			return nil, fmt.Errorf("can't find media for default media id %v", defaultMedia.MediaID)
		}
		dMedia := responses.GetDefaultMedia{}
		dMedia.DefaultID = defaultMedia.ID
		dMedia.ID = media.ID
		dMedia.Key = media.Key
		dMedia.Name = media.Name
		resp.Media = append(resp.Media, dMedia)
	}
	return &resp, nil
}

func MakeResponsePostDefaultMedia(id, mediaID int) *responses.PostDefaultMedia {
	return &responses.PostDefaultMedia{
		ID:      id,
		MediaId: mediaID,
	}
}
