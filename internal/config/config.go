package config

import (
	"path/filepath"
	"strings"

	"github.com/kkyr/fig"
	"github.com/pkg/errors"
)

type AppConfig struct {
	Db     DbConfig     `fig:"db"`
	Server ServerConfig `fig:"server"`
}

func (c *AppConfig) LoadEnv() error {
	cfg, loadErr := loadEnv[AppConfig]("")
	if loadErr != nil {
		return loadErr
	}

	*c = cfg

	return nil
}

// ReadConfig reads configuration data from a file found
// at the passed configPath, returning a populated AppConfig (or error).
func ReadConfig(configPath string) (cfg *AppConfig, err error) {
	var newCfg AppConfig

	var figOpts []fig.Option

	if configPath != "" {
		configPath, err = filepath.Abs(configPath)
		if err != nil {
			return
		}

		fileName := filepath.Base(configPath)
		dirName := filepath.Dir(configPath)

		figOpts = []fig.Option{fig.File(fileName), fig.Dirs(dirName)}
	} else {
		figOpts = []fig.Option{fig.IgnoreFile()}
	}

	err = fig.Load(&newCfg, figOpts...)

	if err != nil {
		err = errors.WithMessagef(err, "error reading from file %[1]s", configPath)
		return
	}

	cfg = &newCfg

	return
}

func loadEnv[T any](envPrefix string) (cfg T, err error) {
	figOpts := []fig.Option{fig.IgnoreFile()}

	if envPrefix != "" {
		figOpts = append(figOpts, fig.UseEnv(envPrefix))
	}

	err = fig.Load(&cfg, figOpts...)

	if err != nil {
		err = errors.WithMessagef(
			err,
			"error loading %[1]s config from environment",
			strings.ToLower(envPrefix),
		)
	}

	return
}

// checkVal checks if a value is equal to its type's zero value as well as a secondary value.
// Returns true if val is not its type's zero value AND val != check.
func checkVal[T comparable](val, check T) bool {
	var zeroVal T

	return val != zeroVal && val != check
}
