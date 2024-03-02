package http

import (
	grpc "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/handler"
)

type GuardHandler struct {
	r           handler.Renderer
	guardClient grpc.GuardClient
}

func NewGuardHandler() {
	return
}
