package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type documentCategoriesServiceStorage interface {
	GetAllDocumentCategories(ctx context.Context) ([]domain.DocumentCategory, error)
	GetDocumentCategory(ctx context.Context, id int) (domain.DocumentCategory, error)
	PostDocumentCategory(ctx context.Context, cat *domain.DocumentCategory) error
	DeleteDocumentCategory(ctx context.Context, id int) error
	UpdateDocumentCategory(ctx context.Context, cat *domain.DocumentCategory) error
}

type DocumentCategoriesService struct {
	storage documentCategoriesServiceStorage
}

func NewDocumentCategoriesService(storage documentCategoriesServiceStorage) *DocumentCategoriesService {
	return &DocumentCategoriesService{storage: storage}
}

func (s *DocumentCategoriesService) GetAllDocumentCategories(ctx context.Context) (*responses.GetAllDocumentCategories, error) {
	cats, err := s.storage.GetAllDocumentCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetAllDocumentCategories: %w", err)
	}
	return mapper.MakeResponseAllDocumentCategories(cats)
}

func (s *DocumentCategoriesService) GetDocumentCategory(ctx context.Context, id int) (*responses.GetDocumentCategory, error) {
	cat, err := s.storage.GetDocumentCategory(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetDocumentCategory: %w", err)
	}
	return mapper.MakeResponseDocumentCategory(&cat)
}

func (s *DocumentCategoriesService) PostDocumentCategory(ctx context.Context, cat *domain.DocumentCategory) error {
	err := s.storage.PostDocumentCategory(ctx, cat)
	if err != nil {
		return fmt.Errorf("can't storage.PostDocumentCategory: %w", err)
	}
	return nil
}

func (s *DocumentCategoriesService) DeleteDocumentCategory(ctx context.Context, id int) error {
	err := s.storage.DeleteDocumentCategory(ctx, id)
	if err != nil {
		return fmt.Errorf("can't storage.DeleteDocumentCategory: %w", err)
	}
	return nil
}

func (s *DocumentCategoriesService) UpdateDocumentCategory(ctx context.Context, cat *domain.DocumentCategory) error {
	err := s.storage.UpdateDocumentCategory(ctx, cat)
	if err != nil {
		return fmt.Errorf("can't storage.UpdateDocumentCategory: %w", err)
	}
	return nil
}
