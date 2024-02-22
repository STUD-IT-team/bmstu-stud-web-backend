package appconfig

import "time"

type Servers struct {
	Public HTTPServer `yaml:"public"`
	Tech   HTTPServer `yaml:"tech"`
}

type HTTPServer struct {
	ListenAddr string `yaml:"listen_addr"`
	BasePath   string `yaml:"base_path"`
}

type GRPCServer struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
