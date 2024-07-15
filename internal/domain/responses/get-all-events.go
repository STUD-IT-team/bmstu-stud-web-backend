package responses

import (
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type GetAllEvents struct {
	Event []Event `json:"event"`
}

type Event struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Prompt      string           `json:"prompt"`
	Media       domain.MediaFile `json:"media"`
	Date        time.Time        `json:"date"`
	Approved    bool             `json:"approved"`
	CreatedAt   time.Time        `json:"created_at"`
	CreatedBy   int              `json:"created_by"`
	RegUrl      string           `json:"reg_url"`
	RegOpenDate time.Time        `json:"reg_open_date"`
	FeedbackUrl string           `json:"feedback_url"`
}
