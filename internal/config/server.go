package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	pathToServerCfg = "config.yaml"
)

type Server struct {
	LogLevel  int    `env-default:"0"     yaml:"level"`
	Port      string `env-default:"50052" yaml:"addr"`
	ConfigLog string `env-default:"prod"  yaml:"configlogger"`
	Network   string `env-default:"tcp"   yaml:"network"`
}

func NewServerCfg() (*Server, error) {
	cfg := &Server{}
	if err := cleanenv.ReadConfig(pathToServerCfg, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
