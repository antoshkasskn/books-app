package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
)

var Cfg *Config

func InitConfig() error {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}
	Cfg = &cfg
	return Cfg.validate()
}

type Config struct {
	DbConfig PostgresDbConfig
	AppPort  int `envconfig:"APP_PORT" default:"8080"`
}

func (c *Config) validate() error {
	if c.AppPort <= 0 {
		return errors.New("app port must be greater than zero")
	}
	if err := c.DbConfig.validate(); err != nil {
		return err
	}
	return nil
}
