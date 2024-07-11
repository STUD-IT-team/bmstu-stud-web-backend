package responses

type GetAllDocuments struct {
	Documents []Document `json:"documents"`
}

type Document struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	ClubID int    `json:"club_id"`
}
