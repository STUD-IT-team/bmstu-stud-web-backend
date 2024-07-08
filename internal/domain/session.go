package domain

import (
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/times"
)

type Session struct {
	SessionID int64     `json:"session_id"`
	MemberID  int       `json:"member_id"`
	IsAdmin   bool      `json:"is_admin"`
	ExpireAt  time.Time `json:"expire_at"`
}

func (s *Session) IsExpired() bool {
	return s.ExpireAt.Before(time.Now().In(times.TZMoscow))
}
