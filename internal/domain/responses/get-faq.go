package responses

type GetFAQ struct {
	ID            int    `db:"id"`
	Question      string `db:"question"`
	Answer        string `db:"answer"`
	Category_id   int    `db:"category_id"`
	Club_id       int    `db:"club_id"`
}

// ID            int    `db:"id"`
// Question      string `db:"question"`
// Answer        string `db:"answer"`
// Type_Question string `db:"type_question"`
// Club_id       int    `db:"club_id"`
