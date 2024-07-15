package responses

type GetDocumentsByClubID struct {
	Documents []Document `json:"documents"`
}
