package domain

import (
	"time"
)

type Session struct {
	UserID    string
	ExpireAt  time.Time
	EnteredAt time.Time
}
