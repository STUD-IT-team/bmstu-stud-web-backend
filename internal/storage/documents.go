package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type documentsStorage interface {
	GetAllDocuments(ctx context.Context) ([]domain.Document, error)
	GetDocument(ctx context.Context, id int) (*domain.Document, error)
}

func (s *storage) GetAllDocuments(ctx context.Context) ([]domain.Document, error) {
	return s.postgres.GetAllDocuments(ctx)
}

func (s *storage) GetDocument(ctx context.Context, id int) (domain.Document, error) {
	return s.postgres.GetDocument(ctx, id)
}
