package borough

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/basiciota"
)

//go:generate stringer -type Borough -linecomment

type Borough basiciota.BasicIota

const (
	Manhattan      Borough = iota // Manhattan
	Brooklyn                      // Brooklyn
	Bronx                         // Bronx
	Queens                        // Queens
	StatenIsland                  // Staten Island
	UnknownBorough                // Unknown
)

var validBoroughs = []Borough{
	Manhattan,
	Brooklyn,
	Bronx,
	Queens,
	StatenIsland,
}

func BoroughFromMtaCsv(s string) Borough {
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
		return UnknownBorough
	}
}

func BoroughFromString(s string) Borough {
	return utils.IotaFromString(s, validBoroughs, UnknownBorough)
}

func (b *Borough) Deserialize(data []byte) error {
	*b = utils.DeserializeIota(data, BoroughFromString)
	return nil
}

func (Borough) GormDataType() string {
	return "borough"
}

func (Borough) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return UnknownBorough.GormDataType()
	default:
		return UnknownBorough.GormDataType()
	}
}

func (Borough) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS ENUM (
	'MANHATTAN',
	'BROOKLYN',
	'BRONX',
	'QUEENS',
	'STATEN_ISLAND'
);`, UnknownBorough.GormDataType())
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of Borough.String
// into a Borough variable.
func (b *Borough) Scan(value any) error {
	*b = utils.DbValueToIota(value.(string), validBoroughs, UnknownBorough)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Borough.String, and no error.
func (b Borough) Value() (driver.Value, error) {
	return utils.IotaToDbValue(b), nil
}
