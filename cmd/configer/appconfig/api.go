package appconfig

import (
	"time"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig/vars"
)

const APIAppName = "api"

type APIConfig struct {
	Servers         Servers           `yaml:"servers"`
	Log             vars.LoggerConfig `yaml:"log"`
	ShutdownTimeout time.Duration     `yaml:"shutdown_timeout"`
}

func GetAPIConfig(env vars.Env) *APIConfig {
	return &APIConfig{
		Servers:         serversConfig(APIAppName),
		Log:             vars.Logger(env),
		ShutdownTimeout: 30 * time.Second,
	}
}

func serversConfig(appName string) Servers {
	return Servers{
		Public: serverPublicConfig(appName),
		Tech:   serverTechConfig(),
	}
}

func serverPublicConfig(appName string) HTTPServer {
	return HTTPServer{
		ListenAddr: vars.PublicAddr,
		BasePath:   basePath(appName),
	}
}

func serverTechConfig() HTTPServer {
	return HTTPServer{
		ListenAddr: vars.TechAddr,
	}
}
