package responses

import (
	"time"
)

type GetEventByID struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Approved    bool      `json:"approved"`
	CreatedAt   time.Time `json:"created_at"`
	RegUrl      int       `json:"reg_url"`
	RegOpenDate time.Time `json:"reg_open_date"`
	FeedbackUrl int       `json:"feedback_url"`
}
