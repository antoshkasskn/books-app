package config

import (
	"errors"
	"github.com/jackc/pgx"
)

type PostgresDbConfig struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     uint16 `envconfig:"DB_PORT" default:"5432"`
	Username string `envconfig:"DB_USERNAME" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"1234"`
	Database string `envconfig:"DB_DATABASE" default:"postgres"`
	MaxConn  int    `envconfig:"DB_MAX_CONN" default:"10"`
}

func (c PostgresDbConfig) GetPgxConf() pgx.ConnPoolConfig {
	return pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     c.Host,
			Port:     c.Port,
			User:     c.Username,
			Password: c.Password,
			Database: c.Database,
		},
		MaxConnections: c.MaxConn,
	}
}

func (c PostgresDbConfig) validate() error {
	if c.MaxConn <= 0 {
		return errors.New("max connections must be greater than zero")
	}
	return nil
}
