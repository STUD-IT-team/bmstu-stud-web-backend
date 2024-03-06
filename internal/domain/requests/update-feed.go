package requests

import (
	"github.com/go-chi/chi"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type UpdateFeed struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	RegistrationURL string    `json:"registration_url"`
	CreatedBy       int       `json:"created_by"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (f *UpdateFeed) Bind(req *http.Request) error {
	err := json.NewDecoder(req.Body).Decode(f)
	if err != nil {
		return fmt.Errorf("can't json decoder on UpdateFeed.Bind: %w", err)
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
