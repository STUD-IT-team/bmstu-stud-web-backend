package app

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type documentsServiceStorage interface {
	GetAllDocuments(ctx context.Context) ([]domain.Document, error)
	GetDocument(ctx context.Context, id int) (domain.Document, error)
	GetDocumentsByClubID(ctx context.Context, clubID int) ([]domain.Document, error)
}

type DocumentsService struct {
	bucketName string
	storage    documentsServiceStorage
}

func NewDocumentsService(storage documentsServiceStorage, bucketName string) *DocumentsService {
	return &DocumentsService{bucketName: bucketName, storage: storage}
}

func (s *DocumentsService) GetAllDocuments(ctx context.Context) (*responses.GetAllDocuments, error) {
	docs, err := s.storage.GetAllDocuments(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.MakeResponseAllDocuments(docs, s.bucketName)
}

func (s *DocumentsService) GetDocument(ctx context.Context, id int) (*responses.GetDocument, error) {
	doc, err := s.storage.GetDocument(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.MakeResponseDocument(&doc, s.bucketName)
}

func (s *DocumentsService) GetDocumentsByClubID(ctx context.Context, clubID int) (*responses.GetDocumentsByClubID, error) {
	docs, err := s.storage.GetDocumentsByClubID(ctx, clubID)
	if err != nil {
		return nil, err
	}
	return mapper.MakeResponseDocumentsByClubID(docs, s.bucketName)
}
