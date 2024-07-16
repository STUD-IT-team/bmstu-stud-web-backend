package domain

type FAQ struct {
	ID            int    `db:"id"`
	Question      string `db:"question"`
	Answer        string `db:"answer"`
	Category_id   int    `db:"category_id"`
	Club_id       int    `db:"club_id"`
}

// id serial primary key,
// question text default '',
// answer text default '',
// type_question text default '',
// club_id int not null
