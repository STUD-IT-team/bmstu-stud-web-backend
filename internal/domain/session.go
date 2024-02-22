package domain

import (
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/times"
)

type Session struct {
	SessionID string    `json:"session_id"`
	MemberID  int64     `json:"member_id"`
	ExpireAt  time.Time `json:"expire_at"`
}

func (s *Session) IsExpired() bool {
	return s.ExpireAt.Before(time.Now().In(times.TZMoscow))
}
