package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/cmd/configer/appconfig/vars"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	appSlice, envSlice := mustParseAppsEnvs(logger)

	for _, app := range appSlice {
		for _, env := range envSlice {
			config := GetConfig(env, app)
			if !reflect.ValueOf(config).IsValid() {
				logger.Warnf("unknown config for app: %s, env: %s", app, env)
				continue
			}

			configPath := filepath.Join(".", "infra", "/", env.String())

			if err := createConfigDir(filepath.Join(".", "infra", "/", env.String())); err != nil {
				logger.Error(err)
				continue
			}

			if err := writeConfig(configPath, "application.conf", config); err != nil {
				logger.WithError(err).Errorf("can't write application config")
			}
		}
	}
}

func mustParseAppsEnvs(logger *logrus.Logger) ([]string, []vars.Env) {
	apps := flag.String("apps", "", "")
	envs := flag.String("envs", "", "local|dev|prod")

	flag.Parse()

	if apps == nil || envs == nil {
		logger.Fatalf("empty apps or envs")

		return nil, nil
	}

	appSlice := strings.Split(*apps, ",")
	envSlice, err := vars.ParseEnvs(*envs, ",")
	if err != nil {
		logger.Fatalf("can't parseEnvs: %v", err)
	}

	if len(appSlice) == 0 || len(envSlice) == 0 {
		logger.Fatalf("empty apps or envs")
	}

	return appSlice, envSlice
}

func createConfigDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("can't create dir '%s': %w", path, err)
		}
	}

	return nil
}

func writeConfig(path, fileName string, data interface{}) error {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("%w: can't marshal yaml: %+v", err, data)
	}

	configFile := filepath.Join(path, "/", fileName)
	f, err := os.Create(configFile)
	if err != nil {
		return fmt.Errorf("%w: can't create file: %s", err, configFile)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if _, err = f.Write(bytes); err != nil {
		return fmt.Errorf("%w, can't write to file: %s", err, configFile)
	}

	if err = f.Sync(); err != nil {
		return fmt.Errorf("%w: can't sync file: %s", err, configFile)
	}

	return nil
}
