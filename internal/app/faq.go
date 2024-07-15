package app

import (
	"context"
	"fmt"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type faqServiceStorage interface {
	GetAllFAQ(ctx context.Context) ([]domain.FAQ, error)
	GetFAQ(ctx context.Context, id int) (domain.FAQ, error)
	GetFAQByClubid(ctx context.Context, club_id int) ([]domain.FAQ, error)
	PostFAQ(ctx context.Context, faq domain.FAQ) error
	DeleteFAQ(ctx context.Context, id int) error
	UpdateFAQ(ctx context.Context, faq domain.FAQ) error
	GetMediaFile(id int) (*domain.MediaFile, error)
	GetMediaFiles(ids []int) (map[int]domain.MediaFile, error)
}

type FAQService struct {
	storage faqServiceStorage
}

func NewFAQService(storage faqServiceStorage) *FAQService {
	return &FAQService{storage: storage}
}
func (s *FAQService) GetAllFAQ(ctx context.Context) (map[string][]responses.FAQ, error) {
	var res []domain.FAQ
	var err error

	res, err = s.storage.GetAllFAQ(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetAllFAQ: %w", err)
	}

	return mapper.MakeResponseAllFAQ(res)
}
func (s *FAQService) GetFAQ(ctx context.Context, id int) (*responses.GetFAQ, error) {
	res, err := s.storage.GetFAQ(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't storage.GetFAQ: %w", err)
	}

	return mapper.MakeResponseFAQ(&res)
}

// func (s *FAQService) GetFAQByClubid(
// 	ctx context.Context,
// 	filter requests.GetFAQByClubid,
// ) (*responses.GetFAQByClubid, error) {
// 	var res []domain.FAQ
// 	var err error

// 	res, err = s.storage.GetFAQByClubid(ctx, filter.Search)
// 	if err != nil {
// 		return nil, fmt.Errorf("can't storage.GetFAQByClubid: %w", err)
// 	}

// 	ids := make([]int, 0, len(res))
// 	for _, faq := range res {
// 		ids = append(ids, faq.MediaID)
// 	}

// 	faqMediaFiles, err := s.storage.GetMediaFiles(ids)
// 	if err != nil {
// 		return nil, fmt.Errorf("can't storage.GetFAQMediaFiles: %w", err)
// 	}

// 	return mapper.MakeResponseFAQByClubid(res, faqMediaFiles)
// }

// func (s *FAQService) PostFAQ(ctx context.Context, faq domain.FAQ) error {
// 	err := s.storage.PostFAQ(ctx, faq)
// 	if err != nil {
// 		return fmt.Errorf("can't storage.PostFAQ: %w", err)
// 	}

// 	return nil
// }

func (s *FAQService) DeleteFAQ(ctx context.Context, id int) error {
	if err := s.storage.DeleteFAQ(ctx, id); err != nil {
		return fmt.Errorf("can't storage.DeleteFAQ: %w", err)
	}

	return nil
}

// // func (s *FAQService) UpdateFAQ(ctx context.Context, faq domain.FAQ) error {
// // 	if err := s.storage.UpdateFAQ(ctx, faq); err != nil {
// // 		return fmt.Errorf("can't storage.UpdateFAQ: %w", err)
// // 	}

// 	return nil
// }
