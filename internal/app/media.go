package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type mediaStorage interface {
	PutMediaFile(ctx context.Context, name string, key string, data []byte) (int, error)
	ClearUpMedia(ctx context.Context, logger *logrus.Logger) error
}

type MediaService struct {
	storage mediaStorage
}

func NewMediaService(storage mediaStorage) *MediaService {
	return &MediaService{
		storage: storage,
	}
}

func (s *MediaService) PostMedia(ctx context.Context, req *requests.PostMedia) (*responses.PostMedia, error) {
	id, err := s.storage.PutMediaFile(ctx, req.Name, req.Name, req.Data)
	if err != nil {
		return &responses.PostMedia{}, fmt.Errorf("can't storage.PutMedia: %v", err)
	}

	return mapper.MakeResponsePostMedia(id), err
}

const bcryptCost = 12

func (s *MediaService) PostMediaBcrypt(ctx context.Context, req *requests.PostMedia) (*responses.PostMedia, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(req.Name), bcryptCost)
	if err != nil {
		return &responses.PostMedia{}, fmt.Errorf("can't bcrypt.GenerateFromPassword: %v", err)
	}

	id, err := s.storage.PutMediaFile(ctx, req.Name, string(key), req.Data)
	if err != nil {
		return &responses.PostMedia{}, fmt.Errorf("can't storage.PutMedia: %v", err)
	}

	return mapper.MakeResponsePostMedia(id), fmt.Errorf("can't storage.UploadObject: %v", err)
}

func (s *MediaService) ClearUpMedia(ctx context.Context, logger *logrus.Logger) error {
	return s.storage.ClearUpMedia(ctx, logger)
}
