package cache

import (
	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

type SessionCache cache.ICache[int64, domain.Session]

// NewSessionCache
//
// Создание нового объекта кэша сессий
func NewSessionCache() SessionCache {
	return cache.NewCache[int64, domain.Session]()
}
