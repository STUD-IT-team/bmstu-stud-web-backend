package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type documentsStorage interface {
	GetAllDocuments(ctx context.Context) ([]domain.Document, error)
	GetDocument(ctx context.Context, id int) (*domain.Document, error)
	GetDocumentsByClubID(ctx context.Context, clubID int) ([]domain.Document, error)
	PostDocument(ctx context.Context, document domain.Document) error
	DeleteDocument(ctx context.Context, id int) error
	UpdateDocument(ctx context.Context, document domain.Document) error
}

func (s *storage) GetAllDocuments(ctx context.Context) ([]domain.Document, error) {
	return s.postgres.GetAllDocuments(ctx)
}

func (s *storage) GetDocument(ctx context.Context, id int) (domain.Document, error) {
	return s.postgres.GetDocument(ctx, id)
}

func (s *storage) GetDocumentsByClubID(ctx context.Context, clubID int) ([]domain.Document, error) {
	return s.postgres.GetDocumentsByClubID(ctx, clubID)
}

func (s *storage) PostDocument(ctx context.Context, document domain.Document) error {
	return s.postgres.PostDocument(ctx, document)
}

func (s *storage) DeleteDocument(ctx context.Context, id int) error {
	return s.postgres.DeleteDocument(ctx, id)
}

func (s *storage) UpdateDocument(ctx context.Context, document domain.Document) error {
	return s.postgres.UpdateDocument(ctx, document)
}
