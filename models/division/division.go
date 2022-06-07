package division

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/basiciota"
)

//go:generate stringer -type Division -linecomment

type Division basiciota.BasicIota

const (
	DivisionBMT     Division = iota // BMT
	DivisionIND                     // IND
	DivisionIRT                     // IRT
	DivisionSIR                     // SIR
	UnknownDivision                 // unknown
)

var validDivisions = []Division{
	DivisionBMT,
	DivisionIND,
	DivisionIRT,
	DivisionSIR,
}

func DivisionFromString(s string) Division {
	return utils.IotaFromString(s, validDivisions, UnknownDivision)
}

func (d *Division) Deserialize(data []byte) error {
	*d = utils.DeserializeIota(data, DivisionFromString)
	return nil
}

func (Division) GormDataType() string {
	return "division"
}

func (Division) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return UnknownDivision.GormDataType()
	default:
		return UnknownDivision.GormDataType()
	}
}

func (Division) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS ENUM (
	'BMT',
	'IND',
	'IRT',
	'SIR'
);`, UnknownDivision.GormDataType())
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of BasicIota.String
// into a Division variable.
func (d *Division) Scan(value any) error {
	*d = utils.DbValueToIota(value.(string), validDivisions, UnknownDivision)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Division.String, and no error.
func (d Division) Value() (driver.Value, error) {
	return utils.IotaToDbValue(d), nil
}
