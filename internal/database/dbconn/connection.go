package dbconn

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/internal/env"
)

type dbConnection struct {
	conn *gorm.DB
	once sync.Once
	sem  *semaphore.Weighted
}

var conn = new(dbConnection)

func errEnvKeyNotSet(envKey string) error {
	return errors.Errorf("%[1]s not set in environment", envKey)
}

func buildDsn(dbConfig *config.DbConfig) string {
	return fmt.Sprintf("host=%[1]s port=%[2]d user=%[3]s password=%[4]s dbname=%[5]s sslmode=%[6]s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Database,
		dbConfig.SslMode,
	)
}

func InitConnection(config *config.AppConfig, logger *zap.Logger) {
	conn.once.Do(func() {
		conn.sem = semaphore.NewWeighted(100)

		dbConfig := &config.Db
		if err := dbConfig.LoadEnv(); err != nil {
			logger.Error("error setting db config fields from environment", zap.Error(err))
		}

		switch {
		case dbConfig.Username == "":
			panic(errEnvKeyNotSet(env.DbUsername))
		case dbConfig.Password == "":
			panic(errEnvKeyNotSet(env.DbPassword))
		case dbConfig.Database == "":
			panic(errEnvKeyNotSet(env.DbDatabase))
		}

		dsn := buildDsn(dbConfig)

		gormConf := &gorm.Config{
			PrepareStmt: true,
		}

		db, err := gorm.Open(postgres.Open(dsn), gormConf)
		if err != nil {
			panic(errors.WithMessage(err, "error opening database connection"))
		}

		sqlDb, err := db.DB()
		if err != nil {
			panic(errors.WithMessage(err, "error fetching underlying connection"))
		}

		if err = sqlDb.Ping(); err != nil {
			panic(errors.WithMessage(err, "error pinging database"))
		}

		sqlDb.SetMaxIdleConns(100)
		sqlDb.SetMaxOpenConns(500)
		sqlDb.SetConnMaxLifetime(time.Hour)

		conn.conn = db
	})
}

func Acquire(ctx context.Context) error {
	return conn.sem.Acquire(ctx, 1)
}

func Release() {
	conn.sem.Release(1)
}

func Connection() *gorm.DB {
	return conn.conn
}

func ConnectionContext(ctx context.Context) *gorm.DB {
	return Connection().WithContext(ctx)
}

func Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	if err := Acquire(ctx); err != nil {
		return err
	}

	defer Release()

	return Connection().Connection(fn)
}
