package domain

type MediaFile struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ClubPhoto struct {
	ID        int `json:"id"`
	MediaID   int `json:"media_id"`
	ClubID    int `json:"club_id"`
	RefNumber int `json:"ref_number"`
}

type DefaultMedia struct {
	ID      int `json:"id"`
	MediaID int `json:"media_id"`
}

type Video struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type MainVideo struct {
	Video
	Current bool `json:"current"`
	ClubID  int  `json:"club_id"`
}
