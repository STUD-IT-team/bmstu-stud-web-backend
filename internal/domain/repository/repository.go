package repository

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"

type IFeedRepository interface {
	GetAllFeed() ([]responses.Feed, error)
	GetFeed(id int) (responses.Feed, error)
}
