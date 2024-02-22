package main

import (
	"fmt"
	"net"
	"os"

	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/postgres"
	serverGRPC "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/handlers/grpc"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/storage"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/pkg/logger"
)

func main() {
	cfg := appconfig.MustParseAppConfig[appconfig.GuardConfig]()

	logger := logger.SetupLogger(cfg.Log)

	if _, err := maxprocs.Set(maxprocs.Logger(logger.Infof)); err != nil {
		logger.WithError(err).Errorf("can't set maxprocs")
	}

	postgres, err := postgres.NewPostgres(os.Getenv("DATA_SOURCE"))
	if err != nil {
		logger.Fatal(err)
	}

	sessionCache := cache.NewSessionCache()
	storage := storage.NewStorage(*postgres, sessionCache)

	guardService := app.NewGuardService(logger, storage)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", os.Getenv("GUARD_DN"), cfg.GRPC.Port))
	if err != nil {
		logger.WithError(err).Errorf("server can't listen and serve requests")
	}

	logger.Infof("starting server, listening on %d", cfg.GRPC.Port)

	grpcServer := grpc.NewServer()
	serverGRPC.Register(grpcServer, guardService)

	logger.Infof("starting service Guard")

	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(err)
	}
}
