package responses

type GetAllFeed struct {
	Feed []Feed `json:"feed"`
}

type Feed struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
