package domain

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Prompt      string    `json:"prompt"`
	MediaID     int       `json:"media_id"`
	Date        time.Time `json:"date"`
	Approved    bool      `json:"approved"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	MainOrg     int       `json:"main_org"`
	ClubID      int       `json:"club_id"`
	RegUrl      string    `json:"reg_url"`
	RegOpenDate time.Time `json:"reg_open_date"`
	FeedbackUrl string    `json:"feedback_url"`
}
