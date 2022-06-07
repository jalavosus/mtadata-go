package enums

import (
	"database/sql/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/utils"
)

// StringEnum is a base type from which "basic" string enums,
// ie. those which are just "const => "string"" can be derived.
// StringEnum provides Marshal and Unmarshal implementations for
// encoding/json and ones yaml package of choice.
// StringEnum.Deserialize, StringEnum.String, and StringEnum.CreateDbType
// must be implemented by the deriving type.
type StringEnum string

func ToStringEnum(s string) StringEnum {
	return StringEnum(s)
}

// Deserialize must be implemented by an iota deriving from BasicIota.
// The implementation function must set the pointer value of a
// variable to the value of the passed []byte.
// How this is done is left up to the developer.
func (b *StringEnum) Deserialize(_ []byte) error {
	panic("not implemented")
}

// String must be implemented by an iota deriving from BasicIota.
// Using something like go:generate stringer is probably the best idea here.
func (b StringEnum) String() string {
	return string(b)
}

// CreateDbType must be implemented by an iota type deriving from BasicIota.
// The implementation function must return a string containing an SQL
// CREATE TYPE AS ENUM command.
func (StringEnum) CreateDbType() string {
	panic("not implemented")
}

func (StringEnum) GormDataType() string {
	panic("not implemented")
}

func (StringEnum) GormDBDataType(*gorm.DB, *schema.Field) string {
	panic("not implemented")
}

// MarshalJSON implements json.Marshaler.
// Returns the JSON-encoded value of BasicIota.String.
func (b StringEnum) MarshalJSON() ([]byte, error) {
	return utils.SerializeIota(b, utils.SerializeJson)
}

// UnmarshalJSON implements json.Unmarshaler.
// Unmarshals the JSON-encoded value of BasicIota.String
// into a BasicIota variable.
func (b *StringEnum) UnmarshalJSON(data []byte) error {
	return b.Deserialize(data)
}

// MarshalYAML implements yaml.Marshaler.
// Returns the YAML-encoded value of BasicIota.String.
func (b StringEnum) MarshalYAML() ([]byte, error) {
	return utils.SerializeIota(b, utils.SerializeYaml)
}

// UnmarshalYAML implements yaml.Unmarshaler.
// Unmarshals the YAML-encoded value of BasicIota.String
// into a BasicIota variable.
func (b *StringEnum) UnmarshalYAML(data []byte) error {
	return b.Deserialize(data)
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of BasicIota.String
// into a BasicIota variable.
func (b *StringEnum) Scan(value any) error {
	return b.Deserialize([]byte(value.(string)))
}

// Value implements driver.Valuer.
// Returns the result of BasicIota.String, and no error.
func (b StringEnum) Value() (driver.Value, error) {
	return b.String(), nil
}
