package main

import (
	"fmt"
	"net"
	"os"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/adapters/cache"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/storage"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/postgres"
	serverGRPC "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/handlers/grpc"

	"github.com/sirupsen/logrus"
	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
)

func main() {
	cfg := appconfig.MustParseAppConfig[appconfig.APIConfig]()

	logger := logrus.New()

	lvl, err := logrus.ParseLevel(cfg.Log.Level)
	if err != nil {
		logger.Fatal(err)
	}

	logger.SetLevel(lvl)

	if _, err := maxprocs.Set(maxprocs.Logger(logger.Infof)); err != nil {
		logger.WithError(err).Errorf("can't set maxprocs")
	}

	postgres, err := postgres.NewPostgres(os.Getenv("DATA_SOURCE"))
	if err != nil {
		logger.Fatal(err)
	}

	storage := storage.NewGuardStorage(*postgres)
	sessionCache := cache.NewSessionCache()

	guardService := app.NewGuardService(logger, storage, sessionCache)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("GUARD_DN"), os.Getenv("GUARD_PORT")))
	if err != nil {
		logger.WithError(err).Errorf("server can't listen and serve requests")
	}

	logger.Infof("starting server, listening on")

	grpcServer := grpc.NewServer()
	serverGRPC.Register(grpcServer, guardService)

	logger.Infof("starting service Guard")

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}
