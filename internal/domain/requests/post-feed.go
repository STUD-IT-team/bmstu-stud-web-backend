package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type PostFeed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Approved    bool   `json:"approved"`
	MediaID     int    `json:"media_id"`
	VkPostUrl   string `json:"vk_post_url"`
}

type PostFeedPointer struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Approved    *bool   `json:"approved"`
	MediaID     *int    `json:"media_id"`
	VkPostUrl   *string `json:"vk_post_url"`
}

func (f *PostFeed) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	pf := PostFeedPointer{}

	err := decoder.Decode(&pf)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostFeed.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object on PostFeed.Bind")
	}

	err = pf.validate()
	if err != nil {
		return fmt.Errorf("%v: %v", domain.ErrIncorrectRequest, err)
	}

	*f = PostFeed{
		Title:       *pf.Title,
		Description: *pf.Description,
		Approved:    *pf.Approved,
		MediaID:     *pf.MediaID,
		VkPostUrl:   *pf.VkPostUrl,
	}

	return f.validate()
}

func (f *PostFeed) validate() error {
	return nil
}

func (pf *PostFeedPointer) validate() error {
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
	return nil
}
