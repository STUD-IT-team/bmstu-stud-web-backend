package storage

import (
	"context"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/hasher"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/times"
	"github.com/google/uuid"
)

type guardStorage interface {
	GetMemberAndValidatePassword(ctx context.Context, login string, password string) (domain.Member, error)
	SetSession(id string, value domain.Session)
	FindSession(id string) *domain.Session
	DeleteSession(id string)
	CheckSession(accessToken string) (*domain.Session, error)
	SaveSessoinFromMemberID(memberID int64) (session domain.Session)
}

func (s *storage) GetMemberAndValidatePassword(ctx context.Context, login string, password string) (domain.Member, error) {
	user, err := s.GetMemberByLogin(ctx, login)
	if err != nil {
		return domain.Member{}, err
	}

	err = hasher.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return domain.Member{}, err
	}

	return user, nil
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

func (s *storage) CheckSession(accessToken string) (*domain.Session, error) {
	session := s.FindSession(accessToken)
	if session == nil {
		return &domain.Session{}, domain.ErrNotFound
	}

	if session.IsExpired() {
		s.DeleteSession(session.SessionID)

		return &domain.Session{}, domain.ErrIsExpired
	}

	return session, nil
}

const sessionDuration = 5 * time.Hour

func (s *storage) SaveSessoinFromMemberID(memberID int64) (session domain.Session) {
	sessionID := uuid.NewString()

	session = domain.Session{
		SessionID: sessionID,
		MemberID:  memberID,
		ExpireAt:  time.Now().In(times.TZMoscow).Add(sessionDuration),
	}

	s.sessionCache.Put(sessionID, session)

	return session
}
