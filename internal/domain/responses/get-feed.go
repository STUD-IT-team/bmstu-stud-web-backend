package responses

import (
	"encoding/base64"
	"time"
)

type GetFeed struct {
	ID              int             `json:"id"`
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	RegistrationURL string          `json:"registration_url"`
	Media           base64.Encoding `json:"media"`
	CreatedBy       int             `json:"created_by"`
	UpdatedAt       time.Time       `json:"updated_at"`
}
