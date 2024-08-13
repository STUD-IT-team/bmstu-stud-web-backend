package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/miniostorage"
	"github.com/sirupsen/logrus"
)

var docBucketEnv = "DOCUMENT_BUCKET"

type documentsStorage interface {
	GetAllDocuments(ctx context.Context) ([]domain.Document, error)
	GetDocument(ctx context.Context, id int) (*domain.Document, error)
	GetDocumentsByCategory(ctx context.Context, categoryID int) ([]domain.Document, error)
	GetDocumentsByClubID(ctx context.Context, clubID int) ([]domain.Document, error)
	PostDocument(ctx context.Context, name string, data []byte, clubId, categoryId int) (string, error)
	DeleteDocument(ctx context.Context, id int) error
	UpdateDocument(ctx context.Context, id int, name string, data []byte, clubId, categoryId int) (string, error)
	CleanupDocuments(ctx context.Context, logger *logrus.Logger) error
}

func (s *storage) GetAllDocuments(ctx context.Context) ([]domain.Document, error) {
	var docBucketName = os.Getenv(docBucketEnv)
	if docBucketName == "" {
		return nil, fmt.Errorf("missing %s environment variable", docBucketEnv)
	}

	docs, err := s.postgres.GetAllDocuments(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't postgres.GetAllDocuments: %w", err)
	}

	for _, v := range docs {
		v.Key = fmt.Sprintf("%s/%s", docBucketName, v.Key)
	}
	return docs, nil
}

func (s *storage) GetDocument(ctx context.Context, id int) (*domain.Document, error) {
	var docBucketName = os.Getenv(docBucketEnv)
	if docBucketName == "" {
		return nil, fmt.Errorf("missing %s environment variable", docBucketEnv)
	}

	doc, err := s.postgres.GetDocument(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't postgres.GetDocument: %w", err)
	}

	doc.Key = fmt.Sprintf("%s/%s", docBucketName, doc.Key)
	return doc, nil
}

func (s *storage) GetDocumentsByCategory(ctx context.Context, categoryID int) ([]domain.Document, error) {
	var docBucketName = os.Getenv(docBucketEnv)
	if docBucketName == "" {
		return nil, fmt.Errorf("missing %s environment variable", docBucketEnv)
	}

	docs, err := s.postgres.GetDocumentsByCategory(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("can't postgres.GetDocumentsByCategory: %w", err)
	}

	for _, v := range docs {
		v.Key = fmt.Sprintf("%s/%s", docBucketName, v.Key)
	}
	return docs, nil
}

func (s *storage) GetDocumentsByClubID(ctx context.Context, clubID int) ([]domain.Document, error) {
	var docBucketName = os.Getenv(docBucketEnv)
	if docBucketName == "" {
		return nil, fmt.Errorf("missing %s environment variable", docBucketEnv)
	}

	docs, err := s.postgres.GetDocumentsByClubID(ctx, clubID)
	if err != nil {
		return nil, fmt.Errorf("can't postgres.GetDocumentsByClubID: %w", err)
	}

	for _, v := range docs {
		v.Key = fmt.Sprintf("%s/%s", docBucketName, v.Key)
	}
	return docs, nil
}

func (s *storage) PostDocument(ctx context.Context, name string, data []byte, clubId, categoryId int) (string, error) {
	var docBucketName = os.Getenv(docBucketEnv)
	if docBucketName == "" {
		return "", fmt.Errorf("missing %s environment variable", docBucketEnv)
	}

	key := fmt.Sprintf("%d/%s", categoryId, name)

	_, err := s.minio.UploadObject(ctx, &miniostorage.UploadObject{
		BucketName: docBucketName,
		ObjectName: key,
		Data:       data,
		Size:       int64(len(data)),
	})
	if err != nil {
		return "", fmt.Errorf("can't minio.UploadObject: %w", err)
	}

	err = s.postgres.PostDocument(ctx, name, key, clubId, categoryId)
	if err != nil {
		// spaghetti wrapping
		err = fmt.Errorf("can't postgres.PostDocument: %w", err) // wrap pgerror

		delErr := s.minio.DeleteObject(ctx, &miniostorage.DeleteObject{
			BucketName: docBucketName,
			ObjectName: key,
		})
		if delErr != nil {
			err = fmt.Errorf("%w && minio.DeleteObject: %w", err, delErr) // add minioerror to wrap (if occurs)
		}
		return "", err // return the final error
	}

	return key, nil
}

func (s *storage) DeleteDocument(ctx context.Context, id int) error {
	var docBucketName = os.Getenv(docBucketEnv)
	if docBucketName == "" {
		return fmt.Errorf("missing %s environment variable", docBucketEnv)
	}

	key, err := s.postgres.DeleteDocument(ctx, id)
	if err != nil {
		return fmt.Errorf("can't postgres.DeleteDocument: %w", err)
	}

	err = s.minio.DeleteObject(ctx, &miniostorage.DeleteObject{
		BucketName: docBucketName,
		ObjectName: key,
	})
	if err != nil {
		return fmt.Errorf("can't minio.DeleteObject: %w", err)
	}

	return nil
}

func (s *storage) UpdateDocument(ctx context.Context, id int, name string, data []byte, clubId, categoryId int) (string, error) {
	var docBucketName = os.Getenv(docBucketEnv)
	if docBucketName == "" {
		return "", fmt.Errorf("missing %s environment variable", docBucketEnv)
	}
	key := fmt.Sprintf("%d/%s", categoryId, name)

	_, err := s.minio.UploadObject(ctx, &miniostorage.UploadObject{
		BucketName: docBucketName,
		ObjectName: key,
		Data:       data,
		Size:       int64(len(data)),
	})
	if err != nil {
		return "", fmt.Errorf("can't minio.UploadObject: %w", err)
	}

	oldKey, err := s.postgres.UpdateDocument(ctx, id, name, key, clubId, categoryId)
	if err != nil {
		err = fmt.Errorf("can't postgres.UpdateDocument: %w", err)
		delErr := s.minio.DeleteObject(ctx, &miniostorage.DeleteObject{
			BucketName: docBucketName,
			ObjectName: key,
		})
		if delErr != nil {
			err = fmt.Errorf("%w && minio.DeleteObject: %w", err, delErr)
		}
		return "", err
	}

	if oldKey != key {
		err = s.minio.DeleteObject(ctx, &miniostorage.DeleteObject{
			BucketName: docBucketName,
			ObjectName: oldKey,
		})
		if err != nil {
			return "", fmt.Errorf("can't minio.DeleteObject: %w", err)
		}
	}

	return key, nil
}

func (s *storage) CleanupDocuments(ctx context.Context, logger *logrus.Logger) error {
	var docBucketName = os.Getenv(docBucketEnv)
	if docBucketName == "" {
		return fmt.Errorf("missing %s environment variable", docBucketEnv)
	}

	logger.Infof("Started deleting unknown documents from object storage...")

	keys, err := s.postgres.GetAllDocumentKeys(ctx)
	if err != nil {
		logger.Warnf("Failed to get all document keys: %v", err)
	}
	logger.Infof("Found %d documents in database: %v", len(keys), keys)

	logger.Infof("Trying to get all document names from object storage...")
	objNames, err := s.minio.GetAllObjectNames(ctx, docBucketName)
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
		err := s.minio.DeleteObject(ctx, &miniostorage.DeleteObject{
			BucketName: docBucketName,
			ObjectName: key,
		})
		if err != nil {
			logger.Warnf("Delete unknown document failed from object storage: %v", err)
			return err
		}
	}
	logger.Infof("Delete unknown documents from object storage successful!")

	return nil
}
