package postgres

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/jackc/pgx"
)

const getAllDocumentCategoriesQuery = "SELECT id, name FROM category"

func (p *Postgres) GetAllDocumentCategories(_ context.Context) ([]domain.DocumentCategory, error) {
	var categories []domain.DocumentCategory

	rows, err := p.db.Query(getAllDocumentCategoriesQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category domain.DocumentCategory

		err = rows.Scan(&category.ID, &category.Name)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if len(categories) == 0 {
		return nil, fmt.Errorf("no categories found")
	}

	return categories, nil
}

const getDocumentCategoryQuery = "SELECT id, name FROM category WHERE id=$1"

func (p *Postgres) GetDocumentCategory(_ context.Context, id int) (*domain.DocumentCategory, error) {
	var category domain.DocumentCategory

	err := p.db.QueryRow(getDocumentCategoryQuery, id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

const postDocumentCategoryQuery = "INSERT INTO category (name) VALUES ($1)"

func (p *Postgres) PostDocumentCategory(_ context.Context, cat *domain.DocumentCategory) error {
	_, err := p.db.Exec(postDocumentCategoryQuery, cat.Name)
	if err != nil {
		return wrapPostgresError(err.(pgx.PgError).Code, err)
	}

	return nil
}

const deleteDocumentCategoryQuery = "DELETE FROM category WHERE id=$1"

func (p *Postgres) DeleteDocumentCategory(_ context.Context, id int) error {
	_, err := p.db.Exec(deleteDocumentCategoryQuery, id)
	if err != nil {
		return wrapPostgresError(err.(pgx.PgError).Code, err)
	}

	return nil
}

const updateDocumentCategoryQuery = "UPDATE category SET name=$1 WHERE id=$2"

func (p *Postgres) UpdateDocumentCategory(_ context.Context, cat *domain.DocumentCategory) error {
	_, err := p.db.Exec(updateDocumentCategoryQuery, cat.Name, cat.ID)
	if err != nil {
		return wrapPostgresError(err.(pgx.PgError).Code, err)
	}

	return nil
}
