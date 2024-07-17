package responses

type GetClearance struct {
	Access  bool   `json:"access"`
	Comment string `json:"comment"`
}
