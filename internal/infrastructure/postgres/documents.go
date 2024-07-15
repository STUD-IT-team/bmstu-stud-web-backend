package postgres

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getAllDocumentsQuery = "SELECT id, name, key, club_id, category_id FROM document"

func (p *Postgres) GetAllDocuments(_ context.Context) ([]domain.Document, error) {
	var documents []domain.Document

	rows, err := p.db.Query(getAllDocumentsQuery)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		var document domain.Document

		err = rows.Scan(&document.ID, &document.Name, &document.Key,
			&document.ClubID, &document.CategoryID)

		if err != nil {
			return nil, wrapPostgresError(err)
		}

		documents = append(documents, document)
	}

	if len(documents) == 0 {
		return nil, ErrPostgresNotFoundError
	}

	return documents, nil
}

const getDocumentsQuery = "SELECT id, name, key, club_id, category_id FROM document WHERE id=$1"

func (p *Postgres) GetDocument(_ context.Context, id int) (*domain.Document, error) {
	var document domain.Document

	err := p.db.QueryRow(getDocumentsQuery, id).Scan(&document.ID, &document.Name,
		&document.Key, &document.ClubID, &document.CategoryID)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	return &document, nil
}

const getDocumentsByCategoryQuery = "SELECT id, name, key, club_id, category_id FROM document WHERE category_id=$1"

func (p *Postgres) GetDocumentsByCategory(_ context.Context, categoryID int) ([]domain.Document, error) {
	var documents []domain.Document

	rows, err := p.db.Query(getDocumentsByCategoryQuery, categoryID)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		var document domain.Document

		err = rows.Scan(&document.ID, &document.Name, &document.Key, &document.ClubID, &document.CategoryID)

		if err != nil {
			return nil, wrapPostgresError(err)
		}

		documents = append(documents, document)
	}

	if len(documents) == 0 {
		return nil, ErrPostgresNotFoundError
	}

	return documents, nil
}

const getDocumentsByClubIDQuery = "SELECT id, name, key, club_id, category_id FROM document WHERE club_id=$1"

func (p *Postgres) GetDocumentsByClubID(_ context.Context, clubID int) ([]domain.Document, error) {
	var documents []domain.Document

	rows, err := p.db.Query(getDocumentsByClubIDQuery, clubID)
	if err != nil {
		return nil, wrapPostgresError(err)
	}

	for rows.Next() {
		var document domain.Document

		err = rows.Scan(&document.ID, &document.Name, &document.Key, &document.ClubID, &document.CategoryID)

		if err != nil {
			return nil, wrapPostgresError(err)
		}

		documents = append(documents, document)
	}

	if len(documents) == 0 {
		return nil, fmt.Errorf("no documents found for club_id=%d", clubID)
	}

	return documents, nil
}

const postDocumentQuery = "INSERT INTO document (name, key, club_id, category_id) VALUES ($1, $2, $3, $4)"

func (p *Postgres) PostDocument(ctx context.Context, name, key string, clubId, categoryId int) error {
	_, err := p.db.Exec(postDocumentQuery, name, key, clubId, categoryId)
	if err != nil {
		return wrapPostgresError(err)
	}

	return nil
}

const getKeyQuery = "SELECT key FROM document WHERE id=$1"
const deleteDocumentQuery = "DELETE FROM document WHERE id=$1"

func (p *Postgres) DeleteDocument(ctx context.Context, id int) (string, error) {
	var key string
	err := p.db.QueryRow(getKeyQuery, id).Scan(&key)
	if err != nil {
		return "", fmt.Errorf("can't get key for document id=%d: %w", id, err)
	}

	tag, err := p.db.Exec(deleteDocumentQuery, id)
	if err != nil {
		return "", wrapPostgresError(err)
	}
	if tag.RowsAffected() == 0 {
		return "", ErrPostgresNotFoundError
	}

	return key, nil
}

const updateDocumentQuery = "UPDATE document SET name = $1, key = $2, club_id = $3, category_id=$4 WHERE id = $5"

func (p *Postgres) UpdateDocument(ctx context.Context, id int, name, key string, clubId, categoryId int) (string, error) {
	var oldKey string
	err := p.db.QueryRow(getKeyQuery, id).Scan(&oldKey)
	if err != nil {
		return "", fmt.Errorf("can't get old key for document id=%d: %w", id, err)
	}

	tag, err := p.db.Exec(updateDocumentQuery, name, key, clubId, categoryId, id)
	if err != nil {
		return "", wrapPostgresError(err)
	}
	if tag.RowsAffected() == 0 {
		return "", ErrPostgresNotFoundError
	}

	return oldKey, nil
}

const getAllDocumentKeysQuery = "SELECT key FROM document"

func (p *Postgres) GetAllDocumentKeys(_ context.Context) ([]string, error) {
	var keys []string

	rows, err := p.db.Query(getAllDocumentKeysQuery)
	if err != nil {
		return nil, fmt.Errorf("can't get all document keys: %w", err)
	}

	for rows.Next() {
		var key string

		err = rows.Scan(&key)
		if err != nil {
			return nil, fmt.Errorf("can't scan row: %w", err)
		}

		keys = append(keys, key)
	}

	if len(keys) == 0 {
		return nil, ErrPostgresNotFoundError
	}

	return keys, nil
}
