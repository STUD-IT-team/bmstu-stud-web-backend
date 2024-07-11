package domain

import "time"

type Event struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Prompt      string    `db:"prompt"`
	MediaID     int       `db:"media_id"`
	Date        time.Time `db:"date"`
	Approved    bool      `db:"approved"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedBy   int       `db:"created_by"`
	RegUrl      string    `db:"reg_url"`
	RegOpenDate time.Time `db:"reg_open_date"`
	FeedbackUrl string    `db:"feedback_url"`
}
