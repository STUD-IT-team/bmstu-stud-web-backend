package responses

type GetAllFeedByTitle struct {
	Feed []Feed `json:"feed"`
}
