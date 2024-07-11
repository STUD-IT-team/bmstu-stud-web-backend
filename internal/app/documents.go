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
	GetDocument(ctx context.Context, id int) (domain.Document, error)
	GetDocumentsByClubID(ctx context.Context, clubID int) ([]domain.Document, error)
	PostDocument(ctx context.Context, name, key string, clubId int) error
	DeleteDocument(ctx context.Context, id int) (string, error)
	UpdateDocument(ctx context.Context, id int, name, key string, clubId int) (string, error)
	GetAllDocumentKeys(ctx context.Context) ([]string, error)
	UploadObject(ctx context.Context, name string, bucketName string, data []byte) (string, error)
	DeleteObject(ctx context.Context, name string, bucketName string) error
	GetAllObjectNames(ctx context.Context, bucketName string) ([]string, error)
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
		return nil, fmt.Errorf("can't storage.GetAllDocuments: %w", err)
	}
	return mapper.MakeResponseAllDocuments(docs, s.bucketName)
}

func (s *DocumentsService) GetDocument(ctx context.Context, id int) (*responses.GetDocument, error) {
	doc, err := s.storage.GetDocument(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetDocument: %w", err)
	}
	return mapper.MakeResponseDocument(&doc, s.bucketName)
}

func (s *DocumentsService) GetDocumentsByClubID(ctx context.Context, clubID int) (*responses.GetDocumentsByClubID, error) {
	docs, err := s.storage.GetDocumentsByClubID(ctx, clubID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetDocumentsByClubID: %w", err)
	}
	return mapper.MakeResponseDocumentsByClubID(docs, s.bucketName)
}

func (s *DocumentsService) PostDocument(ctx context.Context, doc *requests.PostDocument) error {
	_, err := s.storage.UploadObject(ctx, doc.Name, s.bucketName, doc.Data)
	if err != nil {
		return fmt.Errorf("can't storage.UploadObject: %w", err)
	}

	err = s.storage.PostDocument(ctx, doc.Name, doc.Name, doc.ClubID)
	if err != nil {
		s.storage.DeleteObject(ctx, doc.Name, s.bucketName)
		return fmt.Errorf("can't storage.PostDocument: %w", err)
	}

	return nil
}

func (s *DocumentsService) DeleteDocument(ctx context.Context, id int) error {
	name, err := s.storage.DeleteDocument(ctx, id)
	if err != nil {
		return fmt.Errorf("can't storage.DeleteDocument: %w", err)
	}

	err = s.storage.DeleteObject(ctx, name, s.bucketName)
	if err != nil {
		return fmt.Errorf("can't storage.DeleteObject: %w", err)
	}

	return nil
}

func (s *DocumentsService) UpdateDocument(ctx context.Context, doc *requests.UpdateDocument) error {
	_, err := s.storage.UploadObject(ctx, doc.Name, s.bucketName, doc.Data)
	if err != nil {
		return fmt.Errorf("can't storage.UploadObject: %w", err)
	}

	prevName, err := s.storage.UpdateDocument(ctx, doc.ID, doc.Name, doc.Name, doc.ClubID)
	if err != nil {
		s.storage.DeleteObject(ctx, doc.Name, s.bucketName)
		return fmt.Errorf("can't storage.UpdateDocument: %w", err)
	}

	if prevName != doc.Name {
		err = s.storage.DeleteObject(ctx, prevName, s.bucketName)
		if err != nil {
			return fmt.Errorf("can't storage.DeleteObject: %w", err)
		}
	}

	return nil
}

func (s *DocumentsService) ClearUnknownDocuments(ctx context.Context, logger *logrus.Logger) error {
	logger.Infof("Started deleting unknown documents from object storage...")

	keys, err := s.storage.GetAllDocumentKeys(ctx)
	if err != nil {
		logger.Warnf("Failed to get all document keys: %v", err)
	}
	logger.Infof("Found %d documents in database: %v", len(keys), keys)

	logger.Infof("Trying to get all document names from object storage...")
	objNames, err := s.storage.GetAllObjectNames(ctx, s.bucketName)
	if err != nil {
		logger.Warnf("Failed to get all object names: %v", err)
		return err
	}

	logger.Infof("Found %d documents in object storage: %v", len(objNames), objNames)

	unknownKeys := make([]string, 0, len(objNames))

	keysMap := make(map[string]struct{}, len(keys))
	for _, key := range keys {
		keysMap[key] = struct{}{}
	}
	for _, key := range objNames {
		if _, ok := keysMap[key]; !ok {
			unknownKeys = append(unknownKeys, key)
		}
	}

	if len(unknownKeys) == 0 {
		logger.Infof("No unknown documents found")
		return nil
	} else {
		logger.Infof("Found %d unknown documents: %v", len(unknownKeys), unknownKeys)
	}
	logger.Infof("Started deleting unknown documents from object storage...")
	for _, key := range unknownKeys {
		err := s.storage.DeleteObject(ctx, key, s.bucketName)
		if err != nil {
			logger.Warnf("Delete unknown document failed from object storage: %v", err)
			return err
		}
	}
	logger.Infof("Delete unknown documents from object storage successful!")

	return nil
}
