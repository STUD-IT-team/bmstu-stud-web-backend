package storage

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/miniostorage"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/postgres"
	"github.com/sirupsen/logrus"
)

const MEDIA_BUCKET_ENV = "IMAGE_BUCKET"

type mediaFileStorage interface {
	GetMediaFile(_ context.Context, id int) (*domain.MediaFile, error)
	GetMediaFiles(_ context.Context, ids []int) (map[int]domain.MediaFile, error)
	PutMediaFile(ctx context.Context, name string, key string, data []byte) (int, error)
	DeleteUnusedMedia(ctx context.Context, logger *logrus.Logger) error
	DeleteUnknownMedia(ctx context.Context, logger *logrus.Logger) error
	ClearUpMedia(ctx context.Context, logger *logrus.Logger) error
	GetDefaultMedia(ctx context.Context, id int) (*domain.DefaultMedia, error)
	GetAllDefaultMedia(ctx context.Context) ([]domain.DefaultMedia, error)
	PutDefaultMedia(ctx context.Context, name string, key string, data []byte) (id int, mediaId int, err error)
	DeleteDefaultMedia(ctx context.Context, id int) error
	UpdateDefaultMedia(ctx context.Context, id int, name string, data []byte) error
}

func (s *storage) GetMediaFile(_ context.Context, id int) (*domain.MediaFile, error) {
	bucketName := os.Getenv(MEDIA_BUCKET_ENV)
	if bucketName == "" {
		return nil, fmt.Errorf("missing %s environment variable", MEDIA_BUCKET_ENV)
	}

	media, err := s.postgres.GetMediaFile(id)
	if err != nil {
		return nil, fmt.Errorf("can't get media file from postgres: %v", err)
	}
	media.Key = bucketName + "/" + media.Key

	return media, nil
}

func (s *storage) GetMediaFiles(_ context.Context, ids []int) (map[int]domain.MediaFile, error) {
	bucketName := os.Getenv(MEDIA_BUCKET_ENV)
	if bucketName == "" {
		return nil, fmt.Errorf("missing %s environment variable", MEDIA_BUCKET_ENV)
	}

	media, err := s.postgres.GetMediaFiles(ids)
	if err != nil {
		return nil, fmt.Errorf("can't get media file from postgres: %v", err)
	}

	for id, file := range media {
		file.Key = bucketName + "/" + file.Key
		media[id] = file
	}
	return media, nil
}

func (s *storage) PutMediaFile(ctx context.Context, name string, key string, data []byte) (int, error) {
	bucketName := os.Getenv(MEDIA_BUCKET_ENV)
	if bucketName == "" {
		return 0, fmt.Errorf("missing %s environment variable", MEDIA_BUCKET_ENV)
	}

	minioKey, err := s.minio.UploadObject(ctx, &miniostorage.UploadObject{
		BucketName:  bucketName,
		ObjectName:  key,
		Data:        data,
		Size:        int64(len(data)),
		ContentType: "",
	})

	if err != nil {
		return 0, err
	}

	return s.postgres.AddMediaFile(name, minioKey)
}

// Delete media all media from db and object storage if it is not used in other tables.
func (s *storage) DeleteUnusedMedia(ctx context.Context, logger *logrus.Logger) error {
	bucketName := os.Getenv(MEDIA_BUCKET_ENV)
	if bucketName == "" {
		logger.Warnf("missing %s environment variable", MEDIA_BUCKET_ENV)
		return fmt.Errorf("missing %s environment variable", MEDIA_BUCKET_ENV)
	}

	logger.Infof("Started deleting unused media files...")
	med, err := s.postgres.GetUnusedMedia(ctx)

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

	err = s.postgres.DeleteMediaFiles(ctx, keys)
	if err != nil {
		logger.Warnf("Delete unused media failed from database: %v", err)
		return err
	}
	logger.Infof("Delete unused media from database successful!")
	logger.Infof("Started deleting unused media files from object storage...")

	for _, key := range keys {
		req := miniostorage.DeleteObject{
			BucketName: bucketName,
			ObjectName: key,
		}
		err := s.minio.DeleteObject(ctx, &req)
		if err != nil {
			logger.Warnf("Delete unused media failed from object storage: %v", err)
			return err
		}
	}
	logger.Infof("Delete unused media from object storage successful!")

	return nil
}

func (s *storage) DeleteUnknownMedia(ctx context.Context, logger *logrus.Logger) error {
	logger.Infof("Started deleting unknown media files from object storage...")
	bucketName := os.Getenv(MEDIA_BUCKET_ENV)
	if bucketName == "" {
		logger.Warnf("missing %s environment variable", MEDIA_BUCKET_ENV)
		return fmt.Errorf("missing %s environment variable", MEDIA_BUCKET_ENV)
	}

	keys, err := s.postgres.GetAllMediaKeys(ctx)
	if err != nil {
		logger.Warnf("Failed to get all media keys: %v", err)
		return err
	}
	logger.Infof("Found %d media files in database: %v", len(keys), keys)

	logger.Infof("Trying to get all media file names from object storage...")
	objNames, err := s.minio.GetAllObjectNames(ctx, bucketName)
	if err != nil {
		logger.Warnf("Failed to get all object names: %v", err)
		return err
	}

	logger.Infof("Found %d media files in object storage: %v", len(objNames), objNames)

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
		logger.Infof("No unknown media files found")
		return nil
	} else {
		logger.Infof("Found %d unknown media files: %v", len(unknownKeys), unknownKeys)
	}
	logger.Infof("Started deleting unknown media files from object storage...")
	for _, key := range unknownKeys {
		req := miniostorage.DeleteObject{
			BucketName: bucketName,
			ObjectName: key,
		}
		err := s.minio.DeleteObject(ctx, &req)
		if err != nil {
			logger.Warnf("Delete unknown media failed from object storage: %v", err)
			return err
		}
	}
	logger.Infof("Delete unknown media from object storage successful!")

	return nil
}

func (s *storage) ClearUpMedia(ctx context.Context, logger *logrus.Logger) error {
	err := s.DeleteUnusedMedia(ctx, logger)
	if err != nil {
		return err
	}
	return s.DeleteUnknownMedia(ctx, logger)
}

func (s *storage) GetDefaultMedia(ctx context.Context, id int) (*domain.DefaultMedia, error) {
	return s.postgres.GetDefautlMedia(ctx, id)
}

func (s *storage) GetAllDefaultMedia(ctx context.Context) ([]domain.DefaultMedia, error) {
	return s.postgres.GetAllDefaultMedia(ctx)
}

func (s *storage) PutDefaultMedia(ctx context.Context, name string, key string, data []byte) (id int, mediaId int, err error) {
	id, err = s.PutMediaFile(ctx, name, key, data)
	if err != nil {
		return 0, 0, fmt.Errorf("can't put media file: %v", err)
	}
	mediaId, err = s.postgres.AddDefaultMedia(ctx, id)
	if err != nil {
		return 0, 0, fmt.Errorf("can't add default media: %v", err)
	}
	return id, mediaId, nil
}

func (s *storage) DeleteDefaultMedia(ctx context.Context, id int) error {
	return s.postgres.DeleteDefaultMedia(ctx, id)
}

func (s *storage) UpdateMediaFile(ctx context.Context, id int, name string, key string, data []byte) error {
	bucketName := os.Getenv(MEDIA_BUCKET_ENV)
	if bucketName == "" {
		return fmt.Errorf("missing %s environment variable", MEDIA_BUCKET_ENV)
	}

	minioKey, err := s.minio.UploadObject(ctx, &miniostorage.UploadObject{
		BucketName:  bucketName,
		ObjectName:  key,
		Data:        data,
		Size:        int64(len(data)),
		ContentType: "",
	})

	if err != nil {
		return err
	}

	return s.postgres.UpdateMediaFile(id, name, minioKey)
}

func (s *storage) UpdateDefaultMedia(ctx context.Context, id int, name string, key string, data []byte) error {
	bucketName := os.Getenv(MEDIA_BUCKET_ENV)
	if bucketName == "" {
		return fmt.Errorf("missing %s environment variable", MEDIA_BUCKET_ENV)
	}

	mediaID, err := s.PutMediaFile(ctx, name, key, data)
	if errors.Is(err, postgres.ErrPostgresUniqueConstraintViolation) {
		media, err := s.postgres.GetMediaFileByKey(ctx, key)
		if err != nil {
			return fmt.Errorf("can't get media file by key: %v", err)
		}
		return s.UpdateMediaFile(ctx, media.ID, name, key, data)
	}
	if err != nil {
		return fmt.Errorf("can't put media file: %v", err)
	}

	return s.postgres.UpdateDefaultMedia(ctx, id, mediaID)
}
