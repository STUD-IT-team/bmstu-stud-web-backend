package app

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/consts"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app/mapper"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
)

type guardServiceStorage interface {
	DeleteSession(id string)
	SaveSessoinFromMemberID(memberID int64) (session domain.Session)
	GetMemberAndValidatePassword(ctx context.Context, login string, password string) (domain.Member, error)
	CheckSession(accessToken string) (*domain.Session, error)
}

type GuardService struct {
	logger  *logrus.Logger
	storage guardServiceStorage
}

func NewGuardService(log *logrus.Logger, storage guardServiceStorage) *GuardService {
	return &GuardService{
		logger:  log,
		storage: storage,
	}
}

func (s *GuardService) Login(ctx context.Context, req *requests.LoginRequest,
) (res *responses.LoginResponse, err error) {
	const op = "appGuard.Login"

	member, err := s.storage.GetMemberAndValidatePassword(ctx, req.Login, req.Password)
	if err != nil {
		s.logger.WithError(err).Warnf("can't storage.GetUserAndValidatePassword %s", op)
		return nil, fmt.Errorf("can't storage.GetUserAndValidatePassword %s: %w", op, err)
	}

	session := s.storage.SaveSessoinFromMemberID(int64(member.ID))

	s.logger.Infof("user %s logged in successfully", member.Login)

	return mapper.CreateResponseLogin(session.SessionID, session.ExpireAt.Format(consts.GrpcTimeFormat)), nil
}

func (s *GuardService) Logout(_ context.Context, req *requests.LogoutRequest) error {
	s.storage.DeleteSession(req.AccessToken)

	s.logger.Infof("user with session %s uccessfully logged out", req.AccessToken)

	return nil
}

func (s *GuardService) Check(ctx context.Context, req *requests.CheckRequest) (res *responses.CheckResponse, err error) {
	// const op = "appGuard.Check"

	// session, err := s.storage.CheckSession(req.AccessToken)
	// if err != nil {
	// 	s.logger.WithError(err).Warnf("can't storage.CheckSession %s", op)

	// 	return mapper.CreateResponseCheck(false, 0),
	// 		fmt.Errorf("can't storage.CheckSession %s: %w", op, err)
	// }

	// s.logger.Infof("user %d is authorized", session.MemberID)

	// return mapper.CreateResponseCheck(true, session.MemberID), nil

	return nil, nil
}
