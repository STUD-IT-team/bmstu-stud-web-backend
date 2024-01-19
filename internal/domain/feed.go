package domain

type Feed struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	RegistrationURL string `json:"registration_url"`
}
