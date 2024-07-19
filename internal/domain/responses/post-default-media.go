package responses

type PostDefaultMedia struct {
	ID      int    `json:"id"`
	MediaId int    `json:"media_id"`
	Key     string `json:"key"`
}
