package database

type CustomDbTyper interface {
	GormDataType() string
	CreateDbType() string
}
