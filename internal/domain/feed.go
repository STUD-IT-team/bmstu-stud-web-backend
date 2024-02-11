package domain

import "time"

type Feed struct {
	ID              int       `db:"id"`
	Title           string    `db:"title"`
	Description     string    `db:"description"`
	RegistrationURL string    `db:"registration_url"`
	Approved        bool      `db:"approved"`
	MediaUrl        string    `db:"media_url"`
	UpdatedAt       time.Time `db:"updated_at"`
	CreatedAt       time.Time `db:"created_at"`
	CreatedBy       int       `db:"created_by"`
	Views           int       `db:"views"`
}
