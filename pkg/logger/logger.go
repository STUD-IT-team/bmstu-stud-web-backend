package logger

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig/vars"
	"github.com/sirupsen/logrus"
)

func SetupLogger(log vars.LoggerConfig) *logrus.Logger {
	logger := logrus.New()

	lvl, err := logrus.ParseLevel(log.Level)
	if err != nil {
		logger.Fatal(err)
	}

	logger.SetLevel(lvl)

	return logger
}
