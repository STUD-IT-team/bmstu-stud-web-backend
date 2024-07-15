package responses

type GetClubsByType struct {
	Clubs []Club `json:"clubs"`
}
