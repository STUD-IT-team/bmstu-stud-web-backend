package storage

import (
	"context"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type faqStorage interface {
	GetAllFAQ(ctx context.Context) ([]domain.FAQ, error)
	GetFAQByClubid(ctx context.Context, club_id int) ([]domain.FAQ, error)
	GetFAQ(ctx context.Context, id int) (domain.FAQ, error)
	PostFAQ(ctx context.Context, faq domain.FAQ) error
	DeleteFAQ(ctx context.Context, id int) error
	UpdateFAQ(_ context.Context, faq domain.FAQ) error
	// GetFAQByFilter(ctx context.Context, limit, offset int) ([]domain.FAQ, error)
	// GetFAQByFilterIdLastAndOffset(_ context.Context, idLast, offset int) ([]domain.FAQ, error)
}

func (s *storage) GetAllFAQ(ctx context.Context) ([]domain.FAQ, error) {
	return s.postgres.GetAllFAQ(ctx)
}

func (s *storage) GetFAQ(ctx context.Context, id int) (domain.FAQ, error) {
	return s.postgres.GetFAQ(ctx, id)
}

func (s *storage) GetFAQByClubid(ctx context.Context, club_id int) ([]domain.FAQ, error) {
	return s.postgres.GetFAQByClubid(ctx, club_id)
}

func (s *storage) PostFAQ(ctx context.Context, faq domain.FAQ) error {
	return s.postgres.PostFAQ(ctx, faq)
}

func (s *storage) DeleteFAQ(ctx context.Context, id int) error {
	return s.postgres.DeleteFAQ(ctx, id)
}

func (s *storage) UpdateFAQ(ctx context.Context, faq domain.FAQ) error {
	return s.postgres.UpdateFAQ(ctx, faq)
}

// func (s *storage) GetFAQByFilterLimitAndOffset(ctx context.Context, limit, offset int) ([]domain.FAQ, error) {
// 	return s.postgres.GetFAQByFilterLimitAndOffset(ctx, limit, offset)
// }

// func (s *storage) GetFAQByFilterIdLastAndOffset(ctx context.Context, idLast, offset int) ([]domain.FAQ, error) {
// 	return s.postgres.GetFAQByFilterIdLastAndOffset(ctx, idLast, offset)
// }
