package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type documentsStorage interface {
	GetAllDocuments(ctx context.Context) ([]domain.Document, error)
	GetDocument(ctx context.Context, id int) (*domain.Document, error)
	GetDocumentsByClubID(ctx context.Context, clubID int) ([]domain.Document, error)
	PostDocument(ctx context.Context, name, key string, clubId int) error
	DeleteDocument(ctx context.Context, id int) (string, error)
	UpdateDocument(ctx context.Context, id int, name, key string, clubId int) (string, error)
	GetAllDocumentKeys(ctx context.Context) ([]string, error)
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

func (s *storage) PostDocument(ctx context.Context, name, key string, clubId int) error {
	return s.postgres.PostDocument(ctx, name, key, clubId)
}

func (s *storage) DeleteDocument(ctx context.Context, id int) (string, error) {
	return s.postgres.DeleteDocument(ctx, id)
}

func (s *storage) UpdateDocument(ctx context.Context, id int, name, key string, clubId int) (string, error) {
	return s.postgres.UpdateDocument(ctx, id, name, key, clubId)
}

func (s *storage) GetAllDocumentKeys(ctx context.Context) ([]string, error) {
	return s.postgres.GetAllDocumentKeys(ctx)
}
