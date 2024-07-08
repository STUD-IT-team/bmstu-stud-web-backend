package responses

type GetEventsByRange struct {
	Event []Event `json:"event"`
}
