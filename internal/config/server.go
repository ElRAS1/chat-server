package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	pathToServerCfg = "config.yaml"
)

type Server struct {
	LogLevel  int    `yaml:"level" env-default:"0"`
	Port      string `yaml:"addr" env-default:"50052"`
	ConfigLog string `yaml:"configlogger" env-default:"prod"`
	Network   string `yaml:"network" env-default:"tcp"`
}

func NewServerCfg() (*Server, error) {
	cfg := &Server{}
	if err := cleanenv.ReadConfig(pathToServerCfg, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
