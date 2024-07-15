package postgres

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getAllFAQQuery = "SELECT id, question, answer, category_id, club_id FROM faq"

func (p *Postgres) GetAllFAQ(_ context.Context) ([]domain.FAQ, error) {
	var faqs []domain.FAQ

	rows, err := p.db.Query(getAllFAQQuery)
	if err != nil {
		return []domain.FAQ{}, err
	}

	for rows.Next() {
		var faq domain.FAQ

		err = rows.Scan(&faq.ID, &faq.Question, &faq.Answer,
			&faq.Category_id, &faq.Club_id)

		if err != nil {
			return []domain.FAQ{}, err
		}

		faqs = append(faqs, faq)
	}

	if len(faqs) == 0 {
		return []domain.FAQ{}, fmt.Errorf("no faqs found")
	}

	return faqs, nil
}

const getFAQQuery = "SELECT id, question, answer, category_id, club_id FROM faq WHERE id=$1"

func (p *Postgres) GetFAQ(_ context.Context, id int) (domain.FAQ, error) {
	var faq domain.FAQ

	err := p.db.QueryRow(getFAQQuery, id).Scan(&faq.ID, &faq.Question, &faq.Answer, &faq.Category_id, &faq.Club_id)
	if err != nil {
		return domain.FAQ{}, err
	}

	return faq, nil
}

const getFAQByClubidQuery = "SELECT id, question, answer, category_id, club_id FROM faq WHERE club_id $1"

func (p *Postgres) GetFAQByClubid(_ context.Context, club_id int) ([]domain.FAQ, error) {
	var faqs []domain.FAQ

	rows, err := p.db.Query(getFAQByClubidQuery, club_id)
	if err != nil {
		return []domain.FAQ{}, err
	}

	for rows.Next() {
		var faq domain.FAQ

		err = rows.Scan(&faq.ID, &faq.Question, &faq.Answer,
			&faq.Category_id, &faq.Club_id)

		if err != nil {
			return []domain.FAQ{}, err
		}

		faqs = append(faqs, faq)
	}

	if len(faqs) == 0 {
		return []domain.FAQ{}, fmt.Errorf("no faqs found")
	}

	return faqs, nil
}

const postFAQQuery = `INSERT INTO faq (question, answer, category_id, club_id)
        VALUES ($1, $2, $3, $4)`

func (p *Postgres) PostFAQ(_ context.Context, faq domain.FAQ) error {
	_, err := p.db.Exec(postFAQQuery,
		faq.Question,
		faq.Answer,
		faq.Category_id,
		faq.Club_id,
	)

	if err != nil {
		return fmt.Errorf("can't insert faq into postgres %w", err)
	}

	return nil
}

const deleteFAQQuery = "DELETE FROM faq WHERE id=$1"

func (p *Postgres) DeleteFAQ(_ context.Context, id int) error {
	_, err := p.db.Exec(deleteFAQQuery, id)
	if err != nil {
		return fmt.Errorf("can't delete faq on postgres %w", err)
	}

	return nil
}

const updateFAQQuery = `
UPDATE faq SET
question=$1,
answer=$2,
category_id=$3,
club_id=$4 WHERE id=$5`

func (p *Postgres) UpdateFAQ(_ context.Context, faq domain.FAQ) error {
	_, err := p.db.Exec(updateFAQQuery,
		faq.Question,
		faq.Answer,
		faq.Category_id,
		faq.Club_id,
		faq.ID,
	)
	if err != nil {
		return fmt.Errorf("can't update faq on postgres %w", err)
	}

	return nil
}

// const getFAQByFilterLimitAndOffsetQuery = `SELECT id, title, description, created_at, created_by
// 											FROM faq ORDER BY id LIMIT $1 OFFSET $2`

// func (p *Postgres) GetFAQByFilterLimitAndOffset(_ context.Context, limit, offset int) ([]domain.FAQ, error) {
// 	var faqs []domain.FAQ

// 	rows, err := p.db.Query(getFAQByFilterLimitAndOffsetQuery, limit, offset)
// 	if err != nil {
// 		return []domain.FAQ{}, err
// 	}

// 	for rows.Next() {
// 		var faq domain.FAQ

// 		err = rows.Scan(&faq.ID, &faq.Title, &faq.Description, &faq.CreatedAt, &faq.CreatedBy)
// 		if err != nil {
// 			return []domain.FAQ{}, err
// 		}

// 		faqs = append(faqs, faq)
// 	}

// 	return faqs, nil
// }

// const getFAQByFilterIdLastAndOffsetQuery = `SELECT id, title, description, created_at, created_by
// 											FROM faq  WHERE id > $1 ORDER BY id LIMIT $2`

// func (p *Postgres) GetFAQByFilterIdLastAndOffset(_ context.Context, idLast, offset int) ([]domain.FAQ, error) {
// 	var faqs []domain.FAQ

// 	rows, err := p.db.Query(getFAQByFilterIdLastAndOffsetQuery, idLast, offset)
// 	if err != nil {
// 		return []domain.FAQ{}, err
// 	}

// 	for rows.Next() {
// 		var faq domain.FAQ

// 		err = rows.Scan(&faq.ID, &faq.Title, &faq.Description, &faq.CreatedAt, &faq.CreatedBy)
// 		if err != nil {
// 			return []domain.FAQ{}, err
// 		}

// 		faqs = append(faqs, faq)
// 	}

// 	return faqs, nil
// }
