package responses

type GetAllDocuments struct {
	Documents []Document `json:"documents"`
}

type Document struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Key    string `db:"key"`
	ClubID int    `db:"club_id"`
}
