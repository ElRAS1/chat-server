package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const pathToCfg = "internal/config/config.yaml"

type Config struct {
	LogLevel  int    `yaml:"level" env-default:"0"`
	Port      string `yaml:"addr" env-default:"50052"`
	ConfigLog string `yaml:"configlogger" env-default:"prod"`
	Network   string `yaml:"network" env-default:"tcp"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(pathToCfg, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
