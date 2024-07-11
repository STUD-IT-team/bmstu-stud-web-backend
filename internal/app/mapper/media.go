package mapper

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"

func MakeResponsePostMedia(id int) *responses.PostMedia {
	return &responses.PostMedia{
		ID: id,
	}
}
