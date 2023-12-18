package responses

type GetFeed struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	RegistationURL string `json:"registration_url"`
}
