package main

import (
	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig/vars"
)

func GetConfig(env vars.Env, appName string) any {
	switch appName {
	case appconfig.APIAppName:
		return appconfig.GetAPIConfig(env)
	}

	return nil
}
