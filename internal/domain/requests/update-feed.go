package requests

import (
	"encoding/json"
	"strconv"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/go-chi/chi"

	"fmt"
	"net/http"
)

type UpdateFeed struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Approved    bool   `json:"approved"`
	MediaID     int    `json:"media_id"`
	VkPostUrl   string `json:"vk_post_url"`
	Views       int    `json:"views"`
}

type UpdateFeedPointer struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Approved    *bool   `json:"approved"`
	MediaID     *int    `json:"media_id"`
	VkPostUrl   *string `json:"vk_post_url"`
	Views       *int    `json:"views"`
}

func (f *UpdateFeed) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := UpdateFeedPointer{}

	err := decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on UpdateFeed.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on UpdateFeed.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v on UpdateFeed.Bind: %v", domain.ErrIncorrectRequest, err)
	}

	*f = UpdateFeed{
		Title:       *pf.Title,
		Description: *pf.Description,
		Approved:    *pf.Approved,
		MediaID:     *pf.MediaID,
		VkPostUrl:   *pf.VkPostUrl,
		Views:       *pf.Views,
	}

	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on UpdateFeed.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *UpdateFeed) validate() error {
	if f.ID == 0 {
		return fmt.Errorf("require: id")
	}
	return nil
}

func (pf *UpdateFeedPointer) validate() error {
	if pf.Title == nil {
		return fmt.Errorf("require: Title")
	}
	if pf.Description == nil {
		return fmt.Errorf("require: Description")
	}
	if pf.Approved == nil {
		return fmt.Errorf("require: Approved")
	}
	if pf.MediaID == nil {
		return fmt.Errorf("require: MediaID")
	}
	if pf.VkPostUrl == nil {
		return fmt.Errorf("require: VkPostUrl")
	}
	if pf.Views == nil {
		return fmt.Errorf("require: Views")
	}
	return nil
}
