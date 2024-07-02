package requests

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PostFeed struct {
	ID              int             `json:"id"`
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	RegistrationURL string          `json:"registration_url"`
	Media           base64.Encoding `json:"media"`
	CreatedBy       int             `json:"created_by"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

func (f *PostFeed) Bind(req *http.Request) error {
	err := json.NewDecoder(req.Body).Decode(f)
	if err != nil {
		return fmt.Errorf("can't json decoder on PostFeed.Bind: %w", err)
	}

	return nil
}
