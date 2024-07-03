package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PostFeed struct {
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Approved    bool      `db:"approved"`
	MediaID     int       `db:"media_id"`
	VkPostUrl   string    `db:"vk_post_url"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedBy   int       `db:"created_by"`
	Views       int       `db:"views"`
}

func (f *PostFeed) Bind(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(f)
	if err != nil {
		return fmt.Errorf("can't json decoder on UpdateFeed.Bind: %v", err)
	}

	if decoder.More() {
		return fmt.Errorf("extraneous data after JSON object")
	}

	return nil
}
