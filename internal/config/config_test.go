package config_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/internal/env"
	"github.com/jalavosus/mtadata/internal/utils"
)

type CfgType uint

const (
	DbCfg CfgType = iota
	GrpcCfg
	GatewayCfg
)

const (
	testdataPath   string = "testdata/config"
	dbCfgPath      string = "db"
	grpcCfgPath    string = "grpc"
	gatewayCfgPath string = "gateway"
)

const (
	dbCfgPrefix      string = "db"
	grpcCfgPrefix    string = "grpc"
	gatewayCfgPrefix string = "gateway"
)

const (
	defaultsCfgName string = "defaults"
	customCfgName   string = "custom"
)

var (
	cfgExts = []string{"json", "yaml", "toml"}
)

func buildFilePath(t *testing.T, cfgType CfgType, cfgName string) string {
	t.Helper()

	var cfgPath string

	switch cfgType {
	case DbCfg:
		cfgPath = dbCfgPath
	case GrpcCfg:
		cfgPath = grpcCfgPath
	case GatewayCfg:
		cfgPath = gatewayCfgPath
	}

	return filepath.Join(testdataPath, cfgPath, cfgName)
}

func setOptionalEnv(t *testing.T, key string, val *string) {
	t.Helper()

	if val != nil {
		t.Setenv(key, *val)
	}
}

type DbConfigWant struct {
	Host     string
	Port     int
	SslMode  string
	Username string
	Password string
	Database string
}

type DbConfigEnv struct {
	Host     *string
	Port     *string
	SslMode  *string
	Username *string
	Password *string
	Database *string
}

type ServerConfigWant struct {
	Host string
	Port int
}

type ServerConfigEnv struct {
	Host *string
	Port *string
}

type ReadConfigTestCase[T, U any] struct {
	Name   string
	Env    bool
	Want   T
	SetEnv U
}

func testName(prefix, name string, env bool) string {
	nameParts := []string{name}

	if env {
		nameParts = []string{name, "env"}
	}

	if prefix != "" {
		nameParts = append([]string{prefix}, nameParts...)
	}

	return strings.Join(nameParts, "/")
}

type (
	ReadConfigDbTestCase      ReadConfigTestCase[DbConfigWant, DbConfigEnv]
	ReadConfigGrpcTestCase    ReadConfigTestCase[ServerConfigWant, ServerConfigEnv]
	ReadConfigGatewayTestCase ReadConfigTestCase[ServerConfigWant, ServerConfigEnv]
)

func (tc ReadConfigDbTestCase) TestName() string {
	return testName("", tc.Name, tc.Env)
}

func (tc ReadConfigGrpcTestCase) TestName() string {
	return testName(grpcCfgPrefix, tc.Name, tc.Env)
}

func (tc ReadConfigGatewayTestCase) TestName() string {
	return testName(gatewayCfgPrefix, tc.Name, tc.Env)
}

const (
	testWant     string = "test_"
	wantHost     string = "localhost"
	wantSslMode  string = "ssl_mode"
	wantUsername string = "username"
	wantPassword string = "password"
	wantDatabase string = "database"
)

func TestReadConfig(t *testing.T) {
	t.Run("db", testReadConfigDb)
	t.Run("server", func(t *testing.T) {
		testReadConfigGrpc(t)
		testReadConfigGateway(t)
	})
}

func testReadConfigDb(t *testing.T) {
	testCases := []ReadConfigDbTestCase{
		{
			Name: defaultsCfgName,
			Env:  false,
			Want: DbConfigWant{
				Host:     wantHost,
				Port:     5432,
				SslMode:  "disable",
				Username: wantUsername,
				Password: wantPassword,
				Database: wantDatabase,
			},
		},
		{
			Name: defaultsCfgName,
			Env:  true,
			Want: DbConfigWant{
				Host:     testWant + wantHost,
				Port:     5431,
				SslMode:  testWant + wantSslMode,
				Username: testWant + wantUsername,
				Password: testWant + wantPassword,
				Database: testWant + wantDatabase,
			},
			SetEnv: DbConfigEnv{
				Host:     utils.ToPointer(testWant + wantHost),
				Port:     utils.ToPointer("5431"),
				SslMode:  utils.ToPointer(testWant + wantSslMode),
				Username: utils.ToPointer(testWant + wantUsername),
				Password: utils.ToPointer(testWant + wantPassword),
				Database: utils.ToPointer(testWant + wantDatabase),
			},
		},
		{
			Name: customCfgName,
			Env:  false,
			Want: DbConfigWant{
				Host:    testWant + wantHost,
				Port:    5431,
				SslMode: testWant + wantSslMode,
			},
		},
		{
			Name: customCfgName,
			Env:  true,
			Want: DbConfigWant{
				Host:     testWant + wantHost,
				Port:     5431,
				SslMode:  testWant + wantSslMode,
				Username: testWant + wantUsername,
				Password: testWant + wantPassword,
				Database: testWant + wantDatabase,
			},
			SetEnv: DbConfigEnv{
				Username: utils.ToPointer(testWant + wantUsername),
				Password: utils.ToPointer(testWant + wantPassword),
				Database: utils.ToPointer(testWant + wantDatabase),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName(), func(t *testing.T) {
			for _, ext := range cfgExts {
				t.Run(ext, func(t *testing.T) {
					if tc.Env {
						setOptionalEnv(t, env.DbHost, tc.SetEnv.Host)
						setOptionalEnv(t, env.DbPort, tc.SetEnv.Port)
						setOptionalEnv(t, env.DbSslMode, tc.SetEnv.SslMode)
						setOptionalEnv(t, env.DbUsername, tc.SetEnv.Username)
						setOptionalEnv(t, env.DbPassword, tc.SetEnv.Password)
						setOptionalEnv(t, env.DbDatabase, tc.SetEnv.Database)
					}

					cfgPath := buildFilePath(t, DbCfg, tc.Name) + "." + ext
					cfg, err := config.ReadConfig(cfgPath)

					assert.NoError(t, err)
					assert.NotNil(t, cfg.Db)

					if tc.Env {
						assert.NoError(t, cfg.Db.LoadEnv())
					}

					dbCfg := cfg.Db

					assert.Equal(t, tc.Want.Host, dbCfg.Host)
					assert.Equal(t, tc.Want.Port, dbCfg.Port)
					assert.Equal(t, tc.Want.SslMode, dbCfg.SslMode)
					assert.Equal(t, tc.Want.Username, dbCfg.Username)
					assert.Equal(t, tc.Want.Password, dbCfg.Password)
					assert.Equal(t, tc.Want.Database, dbCfg.Database)
				})
			}
		})
	}
}

func testReadConfigGrpc(t *testing.T) {
	testCases := []ReadConfigGrpcTestCase{
		{
			Name: defaultsCfgName,
			Env:  false,
			Want: ServerConfigWant{
				Host: wantHost,
				Port: 50051,
			},
		},
		{
			Name: defaultsCfgName,
			Env:  true,
			Want: ServerConfigWant{
				Host: testWant + wantHost,
				Port: 50050,
			},
			SetEnv: ServerConfigEnv{
				Host: utils.ToPointer(testWant + wantHost),
				Port: utils.ToPointer("50050"),
			},
		},
		{
			Name: customCfgName,
			Env:  false,
			Want: ServerConfigWant{
				Host: testWant + wantHost,
				Port: 50050,
			},
		},
		{
			Name: customCfgName,
			Env:  true,
			Want: ServerConfigWant{
				Host: testWant + wantHost,
				Port: 50059,
			},
			SetEnv: ServerConfigEnv{
				Port: utils.ToPointer("50059"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName(), func(t *testing.T) {
			for _, ext := range cfgExts {
				t.Run(ext, func(t *testing.T) {
					if tc.Env {
						setOptionalEnv(t, env.GrpcHost, tc.SetEnv.Host)
						setOptionalEnv(t, env.GrpcPort, tc.SetEnv.Port)
					}

					cfgPath := buildFilePath(t, GrpcCfg, tc.Name) + "." + ext
					cfg, err := config.ReadConfig(cfgPath)

					assert.NoError(t, err)
					assert.NotNil(t, cfg.Server)
					assert.NotNil(t, cfg.Server.Grpc)

					if tc.Env {
						assert.NoError(t, cfg.Server.Grpc.LoadEnv())
					}

					serverCfg := cfg.Server.Grpc

					assert.Equal(t, tc.Want.Host, serverCfg.Host)
					assert.Equal(t, tc.Want.Port, serverCfg.Port)
				})
			}
		})
	}
}

func testReadConfigGateway(t *testing.T) {
	testCases := []ReadConfigGatewayTestCase{
		{
			Name: defaultsCfgName,
			Env:  false,
			Want: ServerConfigWant{
				Host: wantHost,
				Port: 9090,
			},
		},
		{
			Name: defaultsCfgName,
			Env:  true,
			Want: ServerConfigWant{
				Host: testWant + wantHost,
				Port: 9089,
			},
			SetEnv: ServerConfigEnv{
				Host: utils.ToPointer(testWant + wantHost),
				Port: utils.ToPointer("9089"),
			},
		},
		{
			Name: customCfgName,
			Env:  false,
			Want: ServerConfigWant{
				Host: testWant + wantHost,
				Port: 9089,
			},
		},
		{
			Name: customCfgName,
			Env:  true,
			Want: ServerConfigWant{
				Host: testWant + wantHost,
				Port: 8081,
			},
			SetEnv: ServerConfigEnv{
				Port: utils.ToPointer("8081"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName(), func(t *testing.T) {
			for _, ext := range cfgExts {
				t.Run(ext, func(t *testing.T) {
					if tc.Env {
						setOptionalEnv(t, env.GatewayHost, tc.SetEnv.Host)
						setOptionalEnv(t, env.GatewayPort, tc.SetEnv.Port)
					}

					cfgPath := buildFilePath(t, GatewayCfg, tc.Name) + "." + ext
					cfg, err := config.ReadConfig(cfgPath)

					assert.NoError(t, err)
					assert.NotNil(t, cfg.Server)
					assert.NotNil(t, cfg.Server.Gateway)

					if tc.Env {
						assert.NoError(t, cfg.Server.Gateway.LoadEnv())
					}

					serverCfg := cfg.Server.Gateway

					assert.Equal(t, tc.Want.Host, serverCfg.Host)
					assert.Equal(t, tc.Want.Port, serverCfg.Port)
				})
			}
		})
	}
}
