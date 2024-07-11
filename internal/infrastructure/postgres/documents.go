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

func (p *Postgres) PostDocument(_ context.Context, document domain.Document) error {
	_, err := p.db.Exec(postDocumentQuery, document.Name, document.Key, document.ClubID)
	if err != nil {
		return err
	}

	return nil
}

const deleteDocumentQuery = "DELETE FROM document WHERE id=$1"

func (p *Postgres) DeleteDocument(_ context.Context, id int) error {
	_, err := p.db.Exec(deleteDocumentQuery, id)
	if err != nil {
		return fmt.Errorf("can't delete document on postgres %w", err)
	}

	return nil
}

const updateDocumentQuery = "UPDATE document SET name=$1, key=$2, club_id=$3 WHERE id=$4"

func (p *Postgres) UpdateDocument(_ context.Context, document domain.Document) error {
	_, err := p.db.Exec(updateDocumentQuery, document.Name, document.Key, document.ClubID, document.ID)
	if err != nil {
		return fmt.Errorf("can't update document on postgres %w", err)
	}

	return nil
}
