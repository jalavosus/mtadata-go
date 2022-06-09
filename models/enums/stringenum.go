package enums

import (
	"database/sql/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/utils"
)

type StringEnum string

func (b *StringEnum) Deserialize(_ []byte) error {
	panic("not implemented")
}

func (b StringEnum) String() string {
	return string(b)
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
	return utils.SerializeEnum(b, utils.SerializeJson)
}

func (b *StringEnum) UnmarshalJSON(data []byte) error {
	return b.Deserialize(data)
}

func (b StringEnum) MarshalYAML() ([]byte, error) {
	return utils.SerializeEnum(b, utils.SerializeYaml)
}

func (b *StringEnum) UnmarshalYAML(data []byte) error {
	return b.Deserialize(data)
}

func (b *StringEnum) Scan(value any) error {
	return b.Deserialize([]byte(value.(string)))
}

func (b StringEnum) Value() (driver.Value, error) {
	return b.String(), nil
}
