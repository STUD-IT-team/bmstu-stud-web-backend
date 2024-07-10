package app

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type mediaStorage interface {
	PutMediaFile(name string, key string) (int, error)
	UploadObject(ctx context.Context, name string, bucketName string, data []byte) (string, error)
	DeleteObject(ctx context.Context, name string, bucketName string) error
	DeleteMediaFile(id int) error
	GetUnusedMedia(ctx context.Context) ([]domain.MediaFile, error)
	DeleteMediaFiles(ctx context.Context, keys []string) error
}

type MediaService struct {
	storage    mediaStorage
	bucketName string
}

func NewMediaService(storage mediaStorage, bucketName string) *MediaService {
	return &MediaService{
		storage:    storage,
		bucketName: bucketName,
	}
}

func (s *MediaService) PostMedia(ctx context.Context, req *requests.PostMedia) (int, error) {
	id, err := s.storage.PutMediaFile(req.Name, req.Name)
	if err != nil {
		return 0, err
	}

	_, err = s.storage.UploadObject(ctx, req.Name, s.bucketName, req.Data)
	if err != nil {
		s.storage.DeleteMediaFile(id)
		return 0, err
	}

	return id, err
}

const bcryptCost = 12

func (s *MediaService) PostMediaBcrypt(ctx context.Context, req *requests.PostMedia) (int, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(req.Name), bcryptCost)
	if err != nil {
		return 0, err
	}

	id, err := s.storage.PutMediaFile(req.Name, string(key))
	if err != nil {
		return 0, err
	}

	_, err = s.storage.UploadObject(ctx, string(key), s.bucketName, req.Data)
	if err != nil {
		s.storage.DeleteMediaFile(id)
		return 0, err
	}

	return id, err
}

// / Delete media file by ID. Not checks if file uses in other tables.
func (s *MediaService) BruteDeleteMedia(ctx context.Context, id int) error {

	// return s.storage.DeleteMediaFile(id)
	return nil
}

// / Delete media all media from db and object storage if it is not used in other tables.
func (s *MediaService) DeleteUnusedMedia(ctx context.Context, logger *logrus.Logger) error {

	logger.Infof("Started deleting unused media files...")
	med, err := s.storage.GetUnusedMedia(ctx)

	if err != nil {
		logger.Warnf("DeleteUnusedMedia failed: %v", err)
		return err
	}

	if len(med) == 0 {
		logger.Infof("No unused media files found")
		return nil
	} else {
		logger.Infof("Found %d unused media files: %v", len(med), med)
	}

	keys := make([]string, 0, len(med))
	for _, m := range med {
		keys = append(keys, m.Key)
	}

	logger.Infof("Started deleting unused media files from database...")

	err = s.storage.DeleteMediaFiles(ctx, keys)
	if err != nil {
		logger.Warnf("Delete unused media failed from database: %v", err)
		return err
	}
	logger.Infof("Delete unused media from database successful!")
	logger.Infof("Started deleting unused media files from object storage...")

	for _, key := range keys {
		err := s.storage.DeleteObject(ctx, key, s.bucketName)
		if err != nil {
			logger.Warnf("Delete unused media failed from object storage: %v", err)
			return err
		}
	}
	logger.Infof("Delete unused media from object storage successful!")

	return nil
}
