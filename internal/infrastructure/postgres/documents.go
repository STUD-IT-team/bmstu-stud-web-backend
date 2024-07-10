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
