package responses

type GetDocumentsByCategory struct {
	Documents []Document `json:"documents"`
}
