package responses

type GetDocument struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Key    string `db:"key"`
	ClubID int    `db:"club_id"`
}
