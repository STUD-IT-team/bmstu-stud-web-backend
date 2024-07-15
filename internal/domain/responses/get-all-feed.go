package responses

import (
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type GetAllFeed struct {
	Feed []Feed `json:"feed"`
}

type Feed struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Approved    bool             `json:"approved"`
	Media       domain.MediaFile `json:"media"`
	VkPostUrl   string           `json:"vk_post_url"`
	UpdatedAt   time.Time        `json:"updated_at"`
	CreatedAt   time.Time        `json:"created_at"`
	CreatedBy   int              `json:"created_by"`
	Views       int              `json:"views"`
}
