package config

import (
	"path/filepath"
	"strings"

	"github.com/kkyr/fig"
	"github.com/pkg/errors"

	"github.com/jalavosus/mtadata/internal/env"
)

type AppConfig struct {
	Db     *DbConfig     `fig:"db"`
	Server *ServerConfig `fig:"server"`
}

func (c *AppConfig) LoadEnv() error {
	cfg, loadErr := loadEnv[AppConfig]("")
	if loadErr != nil {
		return loadErr
	}

	*c = cfg

	return nil
}

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

type ServerConfig struct {
	Grpc    *GrpcServerConfig    `fig:"grpc"`
	Gateway *GatewayServerConfig `fig:"gateway"`
}

func (c *ServerConfig) LoadEnv() error {
	if err := c.Grpc.LoadEnv(); err != nil {
		return err
	}

	if err := c.Gateway.LoadEnv(); err != nil {
		return err
	}

	return nil
}

type GrpcServerConfig struct {
	Host string `fig:"host" default:"localhost"`
	Port int    `fig:"port" default:"50051"`
}

func (c *GrpcServerConfig) LoadEnv() error {
	cfg, loadErr := loadEnv[GrpcServerConfig](env.PrefixGrpc)
	if loadErr != nil {
		return loadErr
	}

	if checkVal(cfg.Host, c.Host) && checkVal(cfg.Host, DefaultHost) {
		c.Host = cfg.Host
	}

	if checkVal(cfg.Port, c.Port) && checkVal(cfg.Port, DefaultPortGrpc) {
		c.Port = cfg.Port
	}

	return nil
}

type GatewayServerConfig struct {
	Host string `fig:"host" default:"localhost"`
	Port int    `fig:"port" default:"9090"`
}

func (c *GatewayServerConfig) LoadEnv() error {
	cfg, loadErr := loadEnv[GatewayServerConfig](env.PrefixGateway)
	if loadErr != nil {
		return loadErr
	}

	if checkVal(cfg.Host, c.Host) && checkVal(cfg.Host, DefaultHost) {
		c.Host = cfg.Host
	}

	if checkVal(cfg.Port, c.Port) && checkVal(cfg.Port, DefaultPortGateway) {
		c.Port = cfg.Port
	}

	return nil
}

type rawConfig struct {
	Db *struct {
		Host     string `default:"localhost"`
		Port     int    `default:"5432"`
		SslMode  string `fig:"ssl_mode" default:"disable"`
		Username string
		Password string
		Database string
	} `fig:"db"`
	Server *struct {
		Grpc *struct {
			Host string `fig:"host" default:"localhost"`
			Port int    `fig:"port" default:"50051"`
		} `fig:"grpc"`
		Gateway *struct {
			Host string `fig:"host" default:"localhost"`
			Port int    `fig:"port" default:"9090"`
		} `fig:"gateway"`
	} `fig:"server"`
}

// ReadConfig reads configuration data from a file found
// at the passed configPath, returning a populated AppConfig (or error).
func ReadConfig(configPath string) (cfg *AppConfig, err error) {
	var newCfg AppConfig

	configPath, err = filepath.Abs(configPath)
	if err != nil {
		return
	}

	fileName := filepath.Base(configPath)
	dirName := filepath.Dir(configPath)

	err = fig.Load(
		&newCfg,
		fig.File(fileName),
		fig.Dirs(dirName),
	)

	if err != nil {
		err = errors.WithMessagef(err, "error reading from file %[1]s", configPath)
		return
	}

	// var (
	// 	dbConf     *DbConfig
	// 	serverConf *ServerConfig
	// )
	//
	// if newCfg.Db != nil {
	// 	var conf DbConfig = *newCfg.Db
	// 	dbConf = &conf
	// }
	//
	// if newCfg.Server != nil {
	// 	serverConf = new(ServerConfig)
	//
	// 	if newCfg.Server.Grpc != nil {
	// 		var conf GrpcServerConfig = *newCfg.Server.Grpc
	// 		serverConf.Grpc = &conf
	// 	}
	//
	// 	if newCfg.Server.Gateway != nil {
	// 		var conf GatewayServerConfig = *newCfg.Server.Gateway
	// 		serverConf.Gateway = &conf
	// 	}
	// }
	//
	// cfg = &AppConfig{
	// 	Db:     dbConf,
	// 	Server: serverConf,
	// }

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
