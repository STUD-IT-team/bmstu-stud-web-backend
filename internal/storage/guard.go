package storage

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/hasher"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/times"
)

var maxSessionID = big.NewInt(math.MaxInt64)

// newSessionID generates a cryptographically random session ID rather than a
// sequential one, since the ID doubles as the bearer access token and must
// not be predictable.
func newSessionID() (int64, error) {
	n, err := rand.Int(rand.Reader, maxSessionID)
	if err != nil {
		return 0, fmt.Errorf("can't generate session id: %w", err)
	}

	return n.Int64(), nil
}

type guardStorage interface {
	GetMemberAndValidatePassword(ctx context.Context, login string, password string) (domain.Member, error)
	SetSession(id int64, value domain.Session)
	FindSession(id int64) (domain.Session, error)
	DeleteSession(id int64)
	CheckSession(accessToken int64) (domain.Session, error)
	CreateSession(memberID int, isAdmin bool) (domain.Session, error)
	RegisterMember(ctx context.Context, member *domain.Member) (int, error)
}

func (s *storage) GetMemberAndValidatePassword(ctx context.Context, login string, password string) (domain.Member, error) {
	user, err := s.GetMemberByLogin(ctx, login)
	if err != nil {
		return domain.Member{}, err
	}

	err = hasher.CompareHashAndPassword(user.HashPassword, []byte(password))
	if err != nil {
		return domain.Member{}, fmt.Errorf("%w, size: %v", err, len(user.HashPassword))
	}

	return *user, nil
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
	var sessionID int64

	free := false

	for cnt := 0; cnt < MaxSessionCreateTries; cnt++ {
		id, err := newSessionID()
		if err != nil {
			return domain.Session{}, err
		}

		if _, err = s.FindSession(id); errors.Is(err, domain.ErrNotFound) {
			sessionID = id
			free = true

			break
		}
	}

	if !free {
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

func (s *storage) RegisterMember(ctx context.Context, member *domain.Member) (int, error) {
	hashPassword, err := hasher.HashPassword(member.Password)
	if err != nil {
		return 0, err
	}
	member.HashPassword = hashPassword
	return s.postgres.AddMember(ctx, member)
}
