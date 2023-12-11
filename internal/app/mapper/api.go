package mapper

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

func CreateEcho() *responses.GetEcho {
	return &responses.GetEcho{
		Message: "echo",
	}
}

func MakeResponseAllFeed(dom []domain.Feed) *[]responses.GetAllFeed {
	res := []responses.GetAllFeed{}
	for _, i := range dom {
		res = append(res, responses.GetAllFeed{
			Id: i.Id, Title: i.Title, Description: i.Description,
		})
	}
	return &res
}

func MakeResponseFeed(dom domain.Feed) *responses.GetFeed {
	return &responses.GetFeed{
		Id: dom.Id, Title: dom.Title, Description: dom.Description, RegistationURL: dom.RegistationURL,
	}
}
