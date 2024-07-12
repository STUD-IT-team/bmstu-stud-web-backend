package storage

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type documentCategoriesStorage interface {
	GetAllDocumentCategories(ctx context.Context) ([]domain.DocumentCategory, error)
	GetDocumentCategory(ctx context.Context, id int) (domain.DocumentCategory, error)
	PostDocumentCategory(ctx context.Context, cat *domain.DocumentCategory) error
	DeleteDocumentCategory(ctx context.Context, id int) error
	UpdateDocumentCategory(ctx context.Context, cat *domain.DocumentCategory) error
}

func (s *storage) GetAllDocumentCategories(ctx context.Context) ([]domain.DocumentCategory, error) {
	return s.postgres.GetAllDocumentCategories(ctx)
}

func (s *storage) GetDocumentCategory(ctx context.Context, id int) (domain.DocumentCategory, error) {
	return s.postgres.GetDocumentCategory(ctx, id)
}

func (s *storage) PostDocumentCategory(ctx context.Context, cat *domain.DocumentCategory) error {
	return s.postgres.PostDocumentCategory(ctx, cat)
}

func (s *storage) DeleteDocumentCategory(ctx context.Context, id int) error {
	return s.postgres.DeleteDocumentCategory(ctx, id)
}

func (s *storage) UpdateDocumentCategory(ctx context.Context, cat *domain.DocumentCategory) error {
	return s.postgres.UpdateDocumentCategory(ctx, cat)
}
