package app

import (
	"github.com/sirupsen/logrus"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type API interface {
	Echo() *responses.GetEcho
}

type apiService struct {
	logger *logrus.Logger
}

func NewAPI(logger *logrus.Logger) API {
	return &apiService{
		logger: logger,
	}
}

func (s *apiService) Echo() *responses.GetEcho {
	return mapper.CreateEcho()
}
