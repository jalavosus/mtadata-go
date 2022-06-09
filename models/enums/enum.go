package enums

import (
	"database/sql/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Enum is a base interface which "basic" string enums,
// ie. those which are just "const => "string"" can implement.
// The following types must be
// Enum.Deserialize, Enum.String, and Enum.CreateDbType
// must be implemented by the deriving type.
type Enum interface {

	// Deserialize sets the pointer value of a
	// variable to the value of the passed []byte.
	// How this is done is left up to the developer.
	Deserialize([]byte) error

	// String returns the `string` representation of
	// an implementing type, and implements fmt.Stringer.
	String() string

	// CreateDbType returns a string containing a valid Postgres
	// CREATE TYPE AS ENUM command.
	CreateDbType() string

	// GormDataType implements gorm's schema.GormDataTypeInterface interface.
	// Returns the datatype name of the deriving type for gorm to use.
	// See https://gorm.io/docs/data_types.html#GormDataTypeInterface for more information.
	GormDataType() string

	// GormDBDataType implements gorm's migrator.GormDataTypeInterface interface.
	// Returns the datatype name of the deriving type based on the database "dialect" being used.
	// For example:
	//
	// 	func (EnumType) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	//			switch db.Dialector.Name() {
	//			case "postgres":
	//				return "enum_type"
	//			case "mysql":
	//				return "TEXT"
	//			default:
	//				// some sort of default
	//			}
	// 	}
	//
	// See https://gorm.io/docs/data_types.html#GormDataTypeInterface for more information.
	GormDBDataType(*gorm.DB, *schema.Field) string

	// MarshalJSON implements json.Marshaler.
	MarshalJSON() ([]byte, error)

	// UnmarshalJSON implements json.Unmarshaler.
	UnmarshalJSON(data []byte) error

	// MarshalYAML implements yaml.Marshaler.
	MarshalYAML() ([]byte, error)

	// UnmarshalYAML implements yaml.Unmarshaler.
	UnmarshalYAML(data []byte) error

	// Scan implements sql.Scanner.
	Scan(value any) error

	// Value implements driver.Valuer.
	Value() (driver.Value, error)
}
