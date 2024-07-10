package domain

type MediaFile struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Key  string `db:"key"`
}

type ClubPhoto struct {
	MediaFile
	ClubID    int `db:"club_id"`
	RefNumber int `db:"ref_number"`
}
