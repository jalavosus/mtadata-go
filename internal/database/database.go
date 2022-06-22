package database

import (
	"go.uber.org/fx"

	"github.com/jalavosus/mtadata/internal/database/dbconn"
)

type CustomDbTyper interface {
	GormDataType() string
	CreateDbType() string
}

var Module = fx.Options(fx.Invoke(dbconn.InitConnection))
