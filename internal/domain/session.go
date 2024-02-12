package domain

import (
	"time"
)

type Session struct {
	SessionID string    `json:"session_id"`
	UserID    string    `json:"user_id"`
	ExpireAt  time.Time `json:"expire_at"`
}

func (s *Session) IsExpired() bool {
	loc, _ := time.LoadLocation("Europe/Moscow")

	return s.ExpireAt.Before(time.Now().In(loc))
}
