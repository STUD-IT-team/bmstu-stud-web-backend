package main

import (
	"net"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/app"
	grpc2 "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/ports/grpc"

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

	guardService := app.NewGuard(logger)

	lis, err := net.Listen("tcp", ":")
	if err != nil {
		logger.WithError(err).Errorf("server can't listen and serve requests")
	}

	logger.Infof("starting server, listening on")

	grpcServer := grpc.NewServer()
	grpc2.RegisterGuardServer(grpcServer, guardService)

	logger.Infof("starting service Guard")

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}
