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
	Unknown = Division("")
)

const (
	gormDataTypePg string = "division"
)

var AllDivisions = []Division{
	BMT,
	IND,
	IRT,
	SIR,
}

func FromString(s string) Division {
	return utils.EnumFromString(s, AllDivisions, Unknown)
}

func (d Division) String() string {
	return string(d)
}

func (d *Division) Deserialize(data []byte) error {
	*d = utils.DeserializeEnum(data, FromString)
	return nil
}

func (Division) GormDataType() string {
	return gormDataTypePg
}

func (Division) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return gormDataTypePg
	default:
		return gormDataTypePg
	}
}

func (Division) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS ENUM (
	'BMT',
	'IND',
	'IRT',
	'SIR'
);`, gormDataTypePg)
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of BasicIota.String
// into a Division variable.
func (d *Division) Scan(value any) error {
	*d = utils.DbValueToEnum(value.(string), AllDivisions, Unknown)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Division.String, and no error.
func (d Division) Value() (driver.Value, error) {
	return utils.EnumToDbValue(d), nil
}

func (d *Division) QueryClause() string {
	return "division = ?"
}

func (d *Division) Arg() *any {
	if d != nil {
		var a any = *d
		return &a
	}

	return nil
}

func (d *Division) Invalid() bool {
	if d == nil {
		return true
	}

	return *d == Unknown
}
