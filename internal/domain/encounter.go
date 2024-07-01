package domain

type Encounter struct {
	ID          string `db:"id"`
	Count       int    `db:"count"`
	Description string `db:"description"`
	ClubID      string `db:"club_id"`
}
