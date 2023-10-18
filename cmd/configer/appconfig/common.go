package appconfig

import "github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig/vars"

func basePath(appName string) string {
	return "/" + vars.Project + "/" + appName + "/"
}
