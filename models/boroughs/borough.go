package boroughs

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/enums"
)

type Borough enums.StringEnum

const (
	Manhattan    Borough = "Manhattan"
	Brooklyn     Borough = "Brooklyn"
	Bronx        Borough = "Bronx"
	Queens       Borough = "Queens"
	StatenIsland Borough = "Staten Island"
	Unknown      Borough = "Unknown"
)

const (
	gormDataTypePostgres string = "borough"
)

var validBoroughs = []Borough{
	Manhattan,
	Brooklyn,
	Bronx,
	Queens,
	StatenIsland,
}

func FromMtaCsvString(s string) Borough {
	switch strings.ToUpper(s) {
	case "M":
		return Manhattan
	case "BK":
		return Brooklyn
	case "BX":
		return Bronx
	case "Q":
		return Queens
	case "SI":
		return StatenIsland
	default:
		return Unknown
	}
}

func FromString(s string) Borough {
	return utils.EnumFromString(s, validBoroughs, Unknown)
}

func (b Borough) String() string {
	return string(b)
}

func (b *Borough) Deserialize(data []byte) error {
	*b = utils.DeserializeEnum(data, FromString)
	return nil
}

func (Borough) GormDataType() string {
	return gormDataTypePostgres
}

func (Borough) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return gormDataTypePostgres
	default:
		return gormDataTypePostgres
	}
}

func (Borough) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS ENUM (
	'MANHATTAN',
	'BROOKLYN',
	'BRONX',
	'QUEENS',
	'STATEN_ISLAND'
);`, gormDataTypePostgres)
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of Borough.String
// into a Borough variable.
func (b *Borough) Scan(value any) error {
	*b = utils.DbValueToEnum(value.(string), validBoroughs, Unknown)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Borough.String, and no error.
func (b Borough) Value() (driver.Value, error) {
	return utils.EnumToDbValue(b), nil
}
