package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	pathToCfg = "config.yaml"
)

type Config struct {
	GRPCPort    string `env-default:"50052"     yaml:"grpc_port"`
	HTTPPort    string `env-default:"8082"      yaml:"http_port"`
	SwaggerPort string `env-default:":8091"     yaml:"swagger_port"`
	Host        string `env-default:"localhost" yaml:"host"`
	Network     string `env-default:"tcp"       yaml:"network"`
	ConfigLog   string `env-default:"prod"      yaml:"config_logger"`
	LogLevel    int    `env-default:"0"         yaml:"level"`
}

func NewServerCfg() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(pathToCfg, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
