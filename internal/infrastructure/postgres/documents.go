package postgres

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

const getAllDocumentsQuery = "SELECT id, name, key, club_id FROM document"

func (p *Postgres) GetAllDocuments(_ context.Context) ([]domain.Document, error) {
	var documentss []domain.Document

	rows, err := p.db.Query(getAllDocumentsQuery)
	if err != nil {
		return []domain.Document{}, err
	}

	for rows.Next() {
		var documents domain.Document

		err = rows.Scan(&documents.ID, &documents.Name, &documents.Key, &documents.ClubID)

		if err != nil {
			return []domain.Document{}, err
		}

		documentss = append(documentss, documents)
	}

	if len(documentss) == 0 {
		return []domain.Document{}, fmt.Errorf("no documents found")
	}

	return documentss, nil
}

const getDocumentsQuery = "SELECT id, name, key, club_id FROM document WHERE id=$1"

func (p *Postgres) GetDocument(_ context.Context, id int) (domain.Document, error) {
	var documents domain.Document

	err := p.db.QueryRow(getDocumentsQuery, id).Scan(&documents.ID, &documents.Name, &documents.Key, &documents.ClubID)
	if err != nil {
		return domain.Document{}, err
	}

	return documents, nil
}

const getDocumentsByClubIDQuery = "SELECT id, name, key, club_id FROM document WHERE club_id=$1"

func (p *Postgres) GetDocumentsByClubID(_ context.Context, clubID int) ([]domain.Document, error) {
	var documentss []domain.Document

	rows, err := p.db.Query(getDocumentsByClubIDQuery, clubID)
	if err != nil {
		return []domain.Document{}, err
	}

	for rows.Next() {
		var documents domain.Document

		err = rows.Scan(&documents.ID, &documents.Name, &documents.Key, &documents.ClubID)

		if err != nil {
			return []domain.Document{}, err
		}

		documentss = append(documentss, documents)
	}

	if len(documentss) == 0 {
		return []domain.Document{}, fmt.Errorf("no documents found for club_id=%d", clubID)
	}

	return documentss, nil
}

const postDocumentQuery = "INSERT INTO document (name, key, club_id) VALUES ($1, $2, $3)"

func (p *Postgres) PostDocument(ctx context.Context, name, key string, clubId int) error {
	_, err := p.db.Exec(postDocumentQuery, name, key, clubId)
	if err != nil {
		return fmt.Errorf("can't post document on postgres %w", err)
	}

	return nil
}

const deleteDocumentQuery = "DELETE FROM document WHERE id=$1 RETURNING name"

func (p *Postgres) DeleteDocument(ctx context.Context, id int) (string, error) {
	var name string
	err := p.db.QueryRow(deleteDocumentQuery, id).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("can't delete document on postgres %w", err)
	}

	return name, nil
}

const getOldKeyQuery = "SELECT key FROM document WHERE id=$1"
const updateDocumentQuery = "UPDATE document SET name = $1, key = $2, club_id = $3 WHERE id = $4"

func (p *Postgres) UpdateDocument(ctx context.Context, id int, name, key string, clubId int) (string, error) {
	var oldKey string
	err := p.db.QueryRow(getOldKeyQuery, id).Scan(&oldKey)
	if err != nil {
		return "", fmt.Errorf("can't get old key for document id=%d: %w", id, err)
	}

	_, err = p.db.Exec(updateDocumentQuery, name, key, clubId, id)
	if err != nil {
		return "", fmt.Errorf("can't update document id=%d: %w", id, err)
	}

	return oldKey, nil
}

const getAllDocumentKeysQuery = "SELECT key FROM document"

func (p *Postgres) GetAllDocumentKeys(_ context.Context) ([]string, error) {
	var keys []string

	rows, err := p.db.Query(getAllDocumentKeysQuery)
	if err != nil {
		return []string{}, fmt.Errorf("can't get all document keys: %w", err)
	}

	for rows.Next() {
		var key string

		err = rows.Scan(&key)
		if err != nil {
			return []string{}, fmt.Errorf("can't scan row: %w", err)
		}

		keys = append(keys, key)
	}

	if len(keys) == 0 {
		return []string{}, fmt.Errorf("no document keys found")
	}

	return keys, nil
}
