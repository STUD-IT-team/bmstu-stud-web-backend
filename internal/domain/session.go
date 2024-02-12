package domain

import (
	"time"
)

type Session struct {
	SessionID string
	UserID    string
	ExpireAt  time.Time
}

func (s *Session) IsExpired() bool {
	loc, _ := time.LoadLocation("Europe/Moscow")

	return s.ExpireAt.Before(time.Now().In(loc))
}
