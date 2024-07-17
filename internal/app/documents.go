package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/sirupsen/logrus"
)

type documentsServiceStorage interface {
	GetAllDocuments(ctx context.Context) ([]domain.Document, error)
	GetDocument(ctx context.Context, id int) (*domain.Document, error)
	GetDocumentsByCategory(ctx context.Context, categoryID int) ([]domain.Document, error)
	GetDocumentsByClubID(ctx context.Context, clubID int) ([]domain.Document, error)
	PostDocument(ctx context.Context, name, key string, data []byte, clubId, categoryId int) error
	DeleteDocument(ctx context.Context, id int) error
	UpdateDocument(ctx context.Context, id int, name, key string, data []byte, clubId, categoryId int) error
	CleanupDocuments(ctx context.Context, logger *logrus.Logger) error
}

type DocumentsService struct {
	storage documentsServiceStorage
}

func NewDocumentsService(storage documentsServiceStorage) *DocumentsService {
	return &DocumentsService{storage: storage}
}

func (s *DocumentsService) GetAllDocuments(ctx context.Context) (*responses.GetAllDocuments, error) {
	docs, err := s.storage.GetAllDocuments(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetAllDocuments: %w", err)
	}
	return mapper.MakeResponseAllDocuments(docs)
}

func (s *DocumentsService) GetDocument(ctx context.Context, id int) (*responses.GetDocument, error) {
	doc, err := s.storage.GetDocument(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetDocument: %w", err)
	}
	return mapper.MakeResponseDocument(doc)
}

func (s *DocumentsService) GetDocumentsByCategory(ctx context.Context, categoryID int) (*responses.GetDocumentsByCategory, error) {
	docs, err := s.storage.GetDocumentsByCategory(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetDocumentsByCategory: %w", err)
	}
	return mapper.MakeResponseDocumentsByCategory(docs)
}

func (s *DocumentsService) GetDocumentsByClubID(ctx context.Context, clubID int) (*responses.GetDocumentsByClubID, error) {
	docs, err := s.storage.GetDocumentsByClubID(ctx, clubID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetDocumentsByClubID: %w", err)
	}
	return mapper.MakeResponseDocumentsByClubID(docs)
}

func (s *DocumentsService) PostDocument(ctx context.Context, doc *requests.PostDocument) (*responses.PostDocument, error) {
	var key = fmt.Sprintf("%d/%s", doc.CategoryID, doc.Name)

	err := s.storage.PostDocument(ctx, doc.Name, key, doc.Data, doc.ClubID, doc.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.PostDocument: %w", err)
	}

	return mapper.MakeResponsePostDocument(key)
}

func (s *DocumentsService) DeleteDocument(ctx context.Context, id int) error {
	err := s.storage.DeleteDocument(ctx, id)
	if err != nil {
		return fmt.Errorf("can't storage.DeleteDocument: %w", err)
	}

	return nil
}

func (s *DocumentsService) UpdateDocument(ctx context.Context, doc *requests.UpdateDocument) (*responses.UpdateDocument, error) {
	var key = fmt.Sprintf("%d/%s", doc.CategoryID, doc.Name)

	err := s.storage.UpdateDocument(ctx, doc.ID, doc.Name, key, doc.Data, doc.ClubID, doc.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.UpdateDocument: %w", err)
	}

	return mapper.MakeResponseUpdateDocument(key)
}

func (s *DocumentsService) CleanupDocuments(ctx context.Context, logger *logrus.Logger) error {
	return s.storage.CleanupDocuments(ctx, logger)
}
