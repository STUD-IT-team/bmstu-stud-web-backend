package responses

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type GetClubMedia struct {
	ID    int         `json:"id"`
	Media []ClubMedia `json:"media"`
}

type ClubMedia struct {
	RefNumber int `json:"ref_number"`
	domain.MediaFile
}
