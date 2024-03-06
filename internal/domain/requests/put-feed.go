package requests

import (
	"github.com/go-chi/chi"

	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type PutFeed struct {
	ID              int             `json:"id"`
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	RegistrationURL string          `json:"registration_url"`
	Media           base64.Encoding `json:"media"`
	CreatedBy       int             `json:"created_by"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

func (f *PutFeed) Bind(req *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		return fmt.Errorf("can't Atoi id on PutFeed.Bind: %w", err)
	}

	f.ID = id

	return f.validate()
}

func (f *PutFeed) validate() error {
	if f.ID == 0 {
		return fmt.Errorf("require: id")
	}

	return nil
}
