package domain

import "time"

type Event struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Propmt      string    `db:"propmt"`
	MediaID     int       `db:"media_id"`
	Date        time.Time `db:"date"`
	Approved    bool      `db:"approved"`
	CreatedAt   time.Time `db:"created_at"`
	RegUrl      int       `db:"reg_url"`
	RegOpenDate time.Time `db:"reg_open_date"`
	FeedbackUrl int       `db:"feedback_url"`
}
