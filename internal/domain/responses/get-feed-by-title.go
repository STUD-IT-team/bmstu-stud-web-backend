package responses

type GetFeedByTitle struct {
	Feed []Feed `json:"feed"`
}
