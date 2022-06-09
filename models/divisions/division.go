package divisions

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/enums"
)

type Division enums.StringEnum

const (
	BMT     = Division("BMT")
	IND     = Division("IND")
	IRT     = Division("IRT")
	SIR     = Division("SIR")
	Unknown = Division("Unknown")
)

const (
	gormDataTypePostgres string = "division"
)

var validDivisions = []Division{
	BMT,
	IND,
	IRT,
	SIR,
}

func FromString(s string) Division {
	return utils.EnumFromString(s, validDivisions, Unknown)
}

func (d Division) String() string {
	return string(d)
}

func (d *Division) Deserialize(data []byte) error {
	*d = utils.DeserializeEnum(data, FromString)
	return nil
}

func (Division) GormDataType() string {
	return gormDataTypePostgres
}

func (Division) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return gormDataTypePostgres
	default:
		return gormDataTypePostgres
	}
}

func (Division) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS ENUM (
	'BMT',
	'IND',
	'IRT',
	'SIR'
);`, gormDataTypePostgres)
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of BasicIota.String
// into a Division variable.
func (d *Division) Scan(value any) error {
	*d = utils.DbValueToEnum(value.(string), validDivisions, Unknown)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Division.String, and no error.
func (d Division) Value() (driver.Value, error) {
	return utils.EnumToDbValue(d), nil
}
