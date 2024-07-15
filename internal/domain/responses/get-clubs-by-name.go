package responses

type GetClubsByName struct {
	Clubs []Club `json:"clubs"`
}
