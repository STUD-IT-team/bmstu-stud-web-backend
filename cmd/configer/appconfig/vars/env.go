package vars

import (
	"fmt"
	"strings"
)

const (
	PublicAddr = ":5000"
	TechAddr   = ":9000"
)

type Env int

const (
	Local Env = iota
	Dev
	Prod
	Unknown
)

func envsName() [7]string {
	return [7]string{
		Local: "local",
		Dev:   "dev",
		Prod:  "prod",
	}
}

func (e Env) String() string {
	return envsName()[int(e)]
}

func parseEnv(e string) Env {
	for i, v := range envsName() {
		if v == e {
			return Env(i)
		}
	}

	return Unknown
}

func ParseEnvs(s, sep string) ([]Env, error) {
	var envs []Env
	for _, e := range strings.Split(s, sep) {
		environment := parseEnv(e)
		if environment == Unknown {
			return nil, fmt.Errorf("unknown env '%s'", e)
		}
		envs = append(envs, environment)
	}

	return envs, nil
}
