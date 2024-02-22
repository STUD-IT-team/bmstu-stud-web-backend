package appconfig

import "github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig/vars"

type GuardConfig struct {
	GRPC GRPCServer        `yaml:"grpc"`
	Log  vars.LoggerConfig `yaml:"log"`
}
