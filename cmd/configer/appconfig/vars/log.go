package vars

import "github.com/sirupsen/logrus"

type LoggerConfig struct {
	Level string
}

func Logger(env Env) LoggerConfig {
	level := logrus.InfoLevel
	if env != Prod {
		level = logrus.DebugLevel
	}

	return LoggerConfig{
		Level: level.String(),
	}
}
