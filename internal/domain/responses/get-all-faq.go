package responses

type GetAllFAQ struct {
	FAQ []FAQ `json:"faq"`
}

type FAQ struct {
	ID            int    `db:"id"`
	Question      string `db:"question"`
	Answer        string `db:"answer"`
	Category_id   string `db:"category_id"`
	Club_id       int    `db:"club_id"`
}
