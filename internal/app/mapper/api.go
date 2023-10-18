package mapper

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"

func CreateEcho() *responses.GetEcho {
	return &responses.GetEcho{
		Message: "echo",
	}
}
