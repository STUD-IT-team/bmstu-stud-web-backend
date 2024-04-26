package storage

import (
	"context"

	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/postgres"
)

type Storage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
	DeleteFeed(ctx context.Context, id int) error
	UpdateFeed(_ context.Context, id int, feed domain.Feed) error
	GetFeedByFilter(ctx context.Context, limit, offset int) ([]domain.Feed, error)
	GetMemberByLogin(ctx context.Context, login string) (domain.Member, error)
	SetSession(id string, value domain.Session)
	FindSession(id string) *domain.Session
	DeleteSession(id string)
	SaveSessoinFromMemberID(memberID int64) (session domain.Session)
	GetMemberAndValidatePassword(ctx context.Context, login string, password string) (domain.Member, error)
	CheckSession(accessToken string) (*domain.Session, error)
	GetFeedByFilterIdLastAndOffset(_ context.Context, idLast, offset int) ([]domain.Feed, error)
}

type storage struct {
	postgres     postgres.Postgres
	sessionCache cache.ICache[string, domain.Session]
}

func NewStorage(postgres postgres.Postgres, sessionCache cache.ICache[string, domain.Session]) *storage {
	return &storage{
		postgres:     postgres,
		sessionCache: sessionCache,
	}
}
