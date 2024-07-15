package responses

type GetFeedEncounters struct {
	Encounters []Encounter `json:"encounter"`
}

type Encounter struct {
	ID          int    `json:"id"`
	Count       string `json:"count"`
	Description string `json:"description"`
	ClubID      int    `json:"club_id"`
}
