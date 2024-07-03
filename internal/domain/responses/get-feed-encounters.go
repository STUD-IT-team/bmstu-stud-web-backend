package responses

type GetFeedEncounters struct {
	Encounters []Encounter `json:"encounter"`
}

type Encounter struct {
	ID          int    `db:"id"`
	Count       string `db:"count"`
	Description string `db:"description"`
	ClubID      int    `db:"club_id"`
}
