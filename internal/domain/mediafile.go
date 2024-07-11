package domain

type MediaFile struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ClubPhoto struct {
	MediaFile
	ClubID    int `json:"club_id"`
	RefNumber int `json:"ref_number"`
}
