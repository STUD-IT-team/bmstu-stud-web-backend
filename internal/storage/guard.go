package storage

import (
	"context"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/hasher"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/times"
)

var nextSessionID int64 = 0

func getNextSessionID() int64 {
	if nextSessionID < 0 {
		nextSessionID = 0
	}
	nextSessionID++
	return nextSessionID
}

type guardStorage interface {
	GetMemberAndValidatePassword(ctx context.Context, login string, password string) (domain.Member, error)
	SetSession(id string, value domain.Session)
	FindSession(id string) *domain.Session
	DeleteSession(id string)
	CheckSession(accessToken string) (*domain.Session, error)
	CreateSession(memberID int, isAdmin bool) (domain.Session, error)
}

func (s *storage) GetMemberAndValidatePassword(ctx context.Context, login string, password string) (domain.Member, error) {
	user, err := s.GetMemberByLogin(ctx, login)
	if err != nil {
		return domain.Member{}, err
	}

	err = hasher.CompareHashAndPassword(user.HashPassword, []byte(password))
	if err != nil {
		return domain.Member{}, err
	}

	return user, nil
}

func (s *storage) SetSession(id int64, value domain.Session) {
	s.sessionCache.Put(id, value)
}

func (s *storage) FindSession(id int64) (domain.Session, error) {
	val := s.sessionCache.Find(id)
	if val == nil {
		return domain.Session{}, domain.ErrNotFound
	}

	return *val, nil
}

func (s *storage) DeleteSession(id int64) {
	s.sessionCache.Delete(id)
}

func (s *storage) CheckSession(accessToken int64) (domain.Session, error) {
	session, err := s.FindSession(accessToken)
	if err != nil {
		return domain.Session{}, domain.ErrNotFound
	}

	if session.IsExpired() {
		s.DeleteSession(session.SessionID)

		return domain.Session{}, ErrIsExpired
	}

	return session, nil
}

const sessionDuration = 5 * time.Hour

const MaxSessionCreateTries = 10

func (s *storage) CreateSession(memberID int, isAdmin bool) (domain.Session, error) {
	sessionID := getNextSessionID()
	cnt := 0
	_, err := s.FindSession(sessionID)
	for err == ErrNotFound && cnt < MaxSessionCreateTries {
		cnt++
		sessionID = getNextSessionID()
		_, err = s.FindSession(sessionID)
	}

	if cnt == MaxSessionCreateTries {
		return domain.Session{}, ErrCantCreateSession
	}

	session := domain.Session{
		SessionID: sessionID,
		MemberID:  memberID,
		ExpireAt:  time.Now().In(times.TZMoscow).Add(sessionDuration),
		IsAdmin:   isAdmin,
	}

	s.sessionCache.Put(sessionID, session)

	return session, nil
}
