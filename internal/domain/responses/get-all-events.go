package responses

import (
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type GetAllEvents struct {
	Event []Event `json:"event"`
}

type Event struct {
	ID          int              `db:"id"`
	Title       string           `db:"title"`
	Description string           `db:"description"`
	Propmt      string           `db:"propmt"`
	Media       domain.MediaFile `db:"media"`
	Date        time.Time        `db:"date"`
	Approved    bool             `db:"approved"`
	CreatedAt   time.Time        `db:"created_at"`
	CreatedBy   int              `db:"created_by"`
	RegUrl      string           `db:"reg_url"`
	RegOpenDate time.Time        `db:"reg_open_date"`
	FeedbackUrl string           `db:"feedback_url"`
}