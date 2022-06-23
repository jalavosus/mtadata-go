package config

import (
	"github.com/jalavosus/mtadata/internal/env"
)

// DbConfig contains configuration data
type DbConfig struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	SslMode  string `fig:"ssl_mode" default:"disable"`
	Username string
	Password string
	Database string
}

func (c *DbConfig) LoadEnv() error {
	cfg, loadErr := loadEnv[DbConfig](env.PrefixDb)
	if loadErr != nil {
		return loadErr
	}

	if checkVal(cfg.Host, c.Host) && checkVal(cfg.Host, DefaultHost) {
		c.Host = cfg.Host
	}

	if checkVal(cfg.Port, c.Port) && checkVal(cfg.Port, DefaultPortDb) {
		c.Port = cfg.Port
	}

	if checkVal(cfg.SslMode, c.SslMode) && checkVal(cfg.SslMode, DefaultSslModeDb) {
		c.SslMode = cfg.SslMode
	}

	if checkVal(cfg.Username, c.Username) {
		c.Username = cfg.Username
	}

	if checkVal(cfg.Password, c.Password) {
		c.Password = cfg.Password
	}

	if checkVal(cfg.Database, c.Database) {
		c.Database = cfg.Database
	}

	return nil
}
