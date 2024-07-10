package app

import (
	"context"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"golang.org/x/crypto/bcrypt"
)

type mediaStorage interface {
	PutMediaFile(name string, key string) (int, error)
	UploadObject(ctx context.Context, name string, bucketName string, data []byte) (string, error)
	DeleteObject(ctx context.Context, name string, bucketName string) error
	DeleteMediaFile(id int) error
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
