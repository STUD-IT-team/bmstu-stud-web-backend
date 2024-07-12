package responses

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type GetDefaultMedia struct {
	domain.MediaFile
	DefaultID int `json:"default_id"`
}
