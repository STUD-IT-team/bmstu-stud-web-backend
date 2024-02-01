package domain

import (
	"time"

	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
)

type Session struct {
	UserID    string
	ExpireAt  time.Time
	EnteredAt time.Time
}

// NewSessionCache
//
// Создание нового объекта кэша сессий
func NewSessionCache() cache.ICache[string, Session] {
	return cache.NewCache[string, Session]()
}
