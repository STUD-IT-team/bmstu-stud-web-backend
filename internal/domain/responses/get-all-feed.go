package responses

import (
	"encoding/base64"
	"time"
)

type GetAllFeed struct {
	Feed []Feed `json:"feed"`
}

type Feed struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Media       base64.Encoding `json:"media"`
	CreatedBy   int             `json:"created_by"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
