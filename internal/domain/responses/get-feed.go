package responses

import (
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type GetFeed struct {
	ID          int              `db:"id"`
	Title       string           `db:"title"`
	Description string           `db:"description"`
	Approved    bool             `db:"approved"`
	Media       domain.MediaFile `db:"media"`
	VkPostUrl   string           `db:"vk_post_url"`
	UpdatedAt   time.Time        `db:"updated_at"`
	CreatedAt   time.Time        `db:"created_at"`
	CreatedBy   int              `db:"created_by"`
	Views       int              `db:"views"`
}
