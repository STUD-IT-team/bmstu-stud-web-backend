package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type mediaStorage interface {
	PutMediaFile(ctx context.Context, name string, key string, data []byte) (int, string, error)
	GetMediaFile(_ context.Context, id int) (*domain.MediaFile, error)
	GetMediaFiles(_ context.Context, ids []int) (map[int]domain.MediaFile, error)
	ClearUpMedia(ctx context.Context, logger *logrus.Logger) error
	GetDefaultMedia(ctx context.Context, id int) (*domain.DefaultMedia, error)
	GetAllDefaultMedia(ctx context.Context) ([]domain.DefaultMedia, error)
	PutDefaultMedia(ctx context.Context, name string, key string, data []byte) (id int, objKey string, mediaId int, err error)
	DeleteDefaultMedia(ctx context.Context, id int) error
	UpdateDefaultMedia(ctx context.Context, id int, name string, key string, data []byte) error
	GetMainVideo(ctx context.Context) (*domain.MainVideo, error)
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
	id, key, err := s.storage.PutMediaFile(ctx, req.Name, req.Name, req.Data)
	if err != nil {
		return &responses.PostMedia{}, fmt.Errorf("can't storage.PutMedia: %v", err)
	}

	return mapper.MakeResponsePostMedia(id, key), nil
}

const bcryptCost = 12

func (s *MediaService) PostMediaBcrypt(ctx context.Context, req *requests.PostMedia) (*responses.PostMedia, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(req.Name), bcryptCost)
	if err != nil {
		return &responses.PostMedia{}, fmt.Errorf("can't bcrypt.GenerateFromPassword: %v", err)
	}

	id, objKey, err := s.storage.PutMediaFile(ctx, req.Name, string(key), req.Data)
	if err != nil {
		return &responses.PostMedia{}, fmt.Errorf("can't storage.PutMedia: %v", err)
	}

	return mapper.MakeResponsePostMedia(id, objKey), nil
}

func (s *MediaService) ClearUpMedia(ctx context.Context, logger *logrus.Logger) error {
	return s.storage.ClearUpMedia(ctx, logger)
}

func (s *MediaService) GetMediaDefault(ctx context.Context, ID int) (*responses.GetDefaultMedia, error) {
	defaultMedia, err := s.storage.GetDefaultMedia(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetDefaultMedia: %w", err)
	}

	media, err := s.storage.GetMediaFile(ctx, defaultMedia.MediaID)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetMediaFile: %w", err)
	}

	return mapper.MakeResponseGetDefaultMedia(defaultMedia, media)
}

func (s *MediaService) GetAllMediaDefault(ctx context.Context) (*responses.GetAllDefaultMedia, error) {
	defaultMedias, err := s.storage.GetAllDefaultMedia(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetAllDefaultMedia: %w", err)
	}

	if len(defaultMedias) == 0 {
		return nil, fmt.Errorf("no default media")
	}
	ids := make([]int, 0, len(defaultMedias))
	for _, defaultMedia := range defaultMedias {
		ids = append(ids, defaultMedia.MediaID)
	}

	mediaFiles, err := s.storage.GetMediaFiles(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetMediaFiles: %w", err)
	}

	return mapper.MakeResponseAllDefaultMedia(defaultMedias, mediaFiles)
}

func (s *MediaService) PutMediaDefault(ctx context.Context, name string, data []byte) (*responses.PostDefaultMedia, error) {
	id, key, mediaID, err := s.storage.PutDefaultMedia(ctx, name, name, data)
	if err != nil {
		return nil, fmt.Errorf("can't storage.PutDefaultMedia: %v", err)
	}
	return mapper.MakeResponsePostDefaultMedia(id, key, mediaID), nil
}

func (s *MediaService) DeleteMediaDefault(ctx context.Context, ID int) error {
	return s.storage.DeleteDefaultMedia(ctx, ID)
}

func (s *MediaService) UpdateMediaDefault(ctx context.Context, ID int, name string, data []byte) error {
	return s.storage.UpdateDefaultMedia(ctx, ID, name, name, data)
}

func (s *MediaService) GetMainVideo(ctx context.Context) (*responses.GetMainVideo, error) {
	vid, err := s.storage.GetMainVideo(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetMainVideo: %w", err)
	}

	return &responses.GetMainVideo{MainVideo: *vid}, nil
}
