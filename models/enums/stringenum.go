package enums

import (
	"database/sql/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type StringEnum uint

func (b StringEnum) String() string {
	panic("not implemented")
}

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
	panic("not implemented")
}

func (b *StringEnum) UnmarshalJSON([]byte) error {
	panic("not implemented")
}

func (b StringEnum) MarshalYAML() ([]byte, error) {
	panic("not implemented")
}

func (b *StringEnum) UnmarshalYAML([]byte) error {
	panic("not implemented")
}

func (b *StringEnum) Scan(any) error {
	panic("not implemented")
}

func (b StringEnum) Value() (driver.Value, error) {
	panic("not implemented")
}
