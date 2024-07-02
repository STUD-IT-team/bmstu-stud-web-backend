package responses

type GetFeedEncounters struct {
	Encounters []Encounter `json:"encounter"`
}

type Encounter struct {
	ID          string `db:"id"`
	Count       string `db:"count"`
	Description string `db:"description"`
	ClubID      string `db:"club_id"`
}
