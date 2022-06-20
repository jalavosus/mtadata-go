package connection

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jalavosus/mtadata/internal/env"
)

const (
	dbHostEnv     string = "DB_HOST"
	dbPortEnv     string = "DB_PORT"
	dbUsernameEnv string = "DB_USERNAME"
	dbPasswordEnv string = "DB_PASSWORD"
	dbDbNameEnv   string = "DB_DBNAME"
)

const (
	defaultPort int    = 5432
	sslMode     string = "disable"
)

var (
	conn     *gorm.DB
	connOnce sync.Once
)

func errEnvKeyNotSet(envKey string) error {
	return errors.Errorf("%[1]s not set in environment", envKey)
}

func buildDsn(host string, port int, username, password, dbName string) string {
	return fmt.Sprintf("host=%[1]s port=%[2]d user=%[3]s password=%[4]s dbname=%[5]s sslmode=%[6]s",
		host,
		port,
		username,
		password,
		dbName,
		sslMode,
	)
}

func newConnection() {
	var (
		dbHost     = env.StringFromEnv(dbHostEnv, "localhost")
		dbPort     = env.IntFromEnv(dbPortEnv, defaultPort)
		dbUsername = env.StringFromEnv(dbUsernameEnv, "")
		dbPassword = env.StringFromEnv(dbPasswordEnv, "")
		dbName     = env.StringFromEnv(dbDbNameEnv, "")
	)

	switch {
	case dbUsername == "":
		panic(errEnvKeyNotSet(dbUsernameEnv))
	case dbPassword == "":
		panic(errEnvKeyNotSet(dbPasswordEnv))
	case dbName == "":
		panic(errEnvKeyNotSet(dbDbNameEnv))
	}

	dsn := buildDsn(dbHost, dbPort, dbUsername, dbPassword, dbName)

	config := &gorm.Config{}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		panic(errors.WithMessage(err, "error opening database connection"))
	}

	conn = db
}

func Connection() *gorm.DB {
	connOnce.Do(newConnection)
	return conn
}

func ConnectionContext(ctx context.Context) *gorm.DB {
	connOnce.Do(newConnection)
	return conn.WithContext(ctx)
}
