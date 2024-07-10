package responses

type GetEventsByRange struct {
	Events []Event `json:"events"`
}
