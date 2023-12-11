package responses

type GetFeed struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	RegistationURL string `json:"registration_url"`
}
