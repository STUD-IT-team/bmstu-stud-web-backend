package appconfig

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

func Bind[T any](conf T, configPath string) error {
	f, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("can't open config file for %T: %w", conf, err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("can't read config for %T from '%s': %w", conf, configPath, err)
	}

	if err = unmarshalAndSetup(bytes, conf); err != nil {
		return fmt.Errorf("can't unmarshal config for %T: %w", conf, err)
	}

	return nil
}

func MustParseAppConfig[T any]() T {
	configFile := flag.String("config-path", "/opt/app/config/application.conf", "Application config file")

	flag.Parse()

	cfg := new(T)
	if err := Bind(cfg, *configFile); err != nil {
		log.Fatalf("can't unmarshal application config: %v", err)
	}

	return *cfg
}

type Trigger interface {
	MustAfter()
}

func unmarshalAndSetup(data []byte, cfg interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: parseHook,
		Result:     cfg,
		TagName:    "yaml",
	})
	if err != nil {
		return err
	}

	var tmpConfig map[string]interface{}
	if err = yaml.Unmarshal(data, &tmpConfig); err != nil {
		return err
	}

	if err = decoder.Decode(tmpConfig); err != nil {
		return err
	}

	if t, ok := cfg.(Trigger); ok {
		t.MustAfter()
	}

	return nil
}

func parseHook(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	if from.Kind() == reflect.String {
		stringData := data.(string) //nolint:errcheck // config package
		switch to {
		case reflect.TypeOf(time.Duration(0)):
			return time.ParseDuration(stringData)
		case reflect.TypeOf(time.Time{}):
			return time.Parse(time.RFC3339, stringData)
		case reflect.TypeOf(""):
			return os.ExpandEnv(stringData), nil
		}
	}
	return data, nil
}
