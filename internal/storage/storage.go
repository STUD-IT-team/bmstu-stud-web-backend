package storage

import (
	"context"
	"time"

	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/postgres"
	"github.com/google/uuid"
)

type Storage interface {
	GetAllFeed(ctx context.Context) ([]domain.Feed, error)
	GetFeed(ctx context.Context, id int) (domain.Feed, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	SetSession(id string, value domain.Session)
	FindSession(id string) *domain.Session
	DeleteSession(id string)
	SaveSessoinFromUserID(userID string) (sessionID string, session domain.Session)
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

func (s *storage) GetAllFeed(ctx context.Context) ([]domain.Feed, error) {
	return s.postgres.GetAllFeed(ctx)
}

func (s *storage) GetFeed(ctx context.Context, id int) (domain.Feed, error) {
	return s.postgres.GetFeed(ctx, id)
}

func (s *storage) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.postgres.GetUserByEmail(ctx, email)
}

func (s *storage) SetSession(id string, value domain.Session) {
	s.sessionCache.Put(id, value)
}

func (s *storage) FindSession(id string) *domain.Session {
	return s.sessionCache.Find(id)
}

func (s *storage) DeleteSession(id string) {
	s.sessionCache.Delete(id)
}

const sessionDuration = 5 * time.Hour

func (s *storage) SaveSessoinFromUserID(userID string) (sessionID string, session domain.Session) {
	sessionID = uuid.NewString()

	loc, _ := time.LoadLocation("Europe/Moscow")

	session = domain.Session{
		UserID:    userID,
		ExpireAt:  time.Now().In(loc).Add(time.Duration(sessionDuration)),
		EnteredAt: time.Now().In(loc),
	}

	s.sessionCache.Put(sessionID, session)

	return sessionID, session
}
