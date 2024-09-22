package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	pathToServerCfg = "internal/config/serverConfig.yaml"
	pathToDbCfg     = "internal/config/dbConfig.yaml"
)

type Server struct {
	LogLevel  int    `yaml:"level" env-default:"0"`
	Port      string `yaml:"addr" env-default:"50052"`
	ConfigLog string `yaml:"configlogger" env-default:"prod"`
	Network   string `yaml:"network" env-default:"tcp"`
}

type Db struct {
	ChatDbName    string `yaml:"chatDbName"`    // Db name
	ChatUsernames string `yaml:"chatUsernames"` // Store array user names
	ChatCreatedAt string `yaml:"chatCreatedAt"` // Time created chat
	ChatUpdatedAt string `yaml:"chatUpdatedAt"` // Time updated chat

	Adapter string
}

func NewServerCfg() (*Server, error) {
	cfg := &Server{}
	if err := cleanenv.ReadConfig(pathToServerCfg, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewDbCfg() (*Db, error) {
	cfg := &Db{}
	if err := cleanenv.ReadConfig(pathToDbCfg, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
