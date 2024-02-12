package storage

import (
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/times"
	"github.com/google/uuid"
)

func (s *storage) SetSession(id string, value domain.Session) {
	s.sessionCache.Put(id, value)
}

func (s *storage) FindSession(id string) *domain.Session {
	return s.sessionCache.Find(id)
}

func (s *storage) DeleteSession(id string) {
	s.sessionCache.Delete(id)
}

func (s *storage) CheckSession(accessToken string) (*domain.Session, error) {
	session := s.FindSession(accessToken)
	if session == nil {
		return &domain.Session{}, domain.ErrNotFound
	}

	if session.IsExpired() {
		s.DeleteSession(session.UserID)

		return &domain.Session{}, domain.ErrIsExpired
	}

	return session, nil
}

const sessionDuration = 5 * time.Hour

func (s *storage) SaveSessoinFromUserID(userID string) (session domain.Session) {
	sessionID := uuid.NewString()

	session = domain.Session{
		SessionID: sessionID,
		UserID:    userID,
		ExpireAt:  time.Now().In(times.TZMoscow).Add(sessionDuration),
	}

	s.sessionCache.Put(sessionID, session)

	return session
}
