package cache

import (
	"github.com/STUD-IT-team/bauman-legends-backend/pkg/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
)

// NewSessionCache
//
// Создание нового объекта кэша сессий
func NewSessionCache() cache.ICache[string, domain.Session] {
	return cache.NewCache[string, domain.Session]()
}
