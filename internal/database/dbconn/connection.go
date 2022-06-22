package dbconn

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/internal/env"
)

var (
	conn     *gorm.DB
	connOnce sync.Once
)

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
	connOnce.Do(func() {
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

		gormConf := &gorm.Config{}

		db, err := gorm.Open(postgres.Open(dsn), gormConf)
		if err != nil {
			panic(errors.WithMessage(err, "error opening database connection"))
		}

		conn = db
	})
}

func Connection() *gorm.DB {
	return conn
}

func ConnectionContext(ctx context.Context) *gorm.DB {
	return Connection().WithContext(ctx)
}
