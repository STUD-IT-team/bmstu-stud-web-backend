package domain

import (
	"time"
)

type Session struct {
	UserID    string
	ExpireAt  time.Time
	EnteredAt time.Time
}

func (s *Session) IsExpired() bool {
	loc, _ := time.LoadLocation("Europe/Moscow")

	return s.ExpireAt.Before(time.Now().In(loc))
}
