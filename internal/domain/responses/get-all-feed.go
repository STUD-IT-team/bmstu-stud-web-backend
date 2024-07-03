package responses

import (
	"time"
)

type GetAllFeed struct {
	Feed []Feed `json:"feed"`
}

type Feed struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Approved    bool      `db:"approved"`
	MediaID     int       `db:"media_id"`
	VkPostUrl   string    `db:"vk_post_url"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedBy   int       `db:"created_by"`
	Views       int       `db:"views"`
}
