package database

import (
	"context"

	"github.com/pkg/errors"

	"github.com/jalavosus/mtadata/internal/database/connection"
)

type CustomDbTyper interface {
	GormDataType() string
	CreateDbType() string
}

func MigrateType(ctx context.Context, dataType any) error {
	conn := connection.Connection()
	conn = conn.WithContext(ctx)

	if customTyper, ok := dataType.(CustomDbTyper); ok {
		cmd := customTyper.CreateDbType()
		if err := conn.Exec(cmd).Error; err != nil {
			return errors.WithMessagef(err, "error executing sql '%[1]s'", cmd)
		}
	} else {
		if err := conn.AutoMigrate(dataType); err != nil {
			return errors.WithMessage(err, "error automigrating type")
		}
	}

	return nil
}
