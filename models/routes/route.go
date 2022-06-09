package routes

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/enums"
)

type Route enums.StringEnum

const (
	One     = Route("1")
	Two     = Route("2")
	Three   = Route("3")
	Four    = Route("4")
	Five    = Route("5")
	Six     = Route("6")
	Seven   = Route("7")
	A       = Route("A")
	B       = Route("B")
	C       = Route("C")
	D       = Route("D")
	E       = Route("E")
	F       = Route("F")
	G       = Route("G")
	J       = Route("J")
	L       = Route("L")
	M       = Route("M")
	N       = Route("N")
	Q       = Route("Q")
	R       = Route("R")
	S       = Route("S")
	SIR     = Route("SIR")
	W       = Route("W")
	Z       = Route("Z")
	Unknown = Route("Unknown")
)

const (
	gormDataTypePostgres string = "route"
)

var validRoutes = []Route{
	One,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	A,
	B,
	C,
	D,
	E,
	F,
	G,
	J,
	L,
	M,
	N,
	Q,
	R,
	S,
	SIR,
	W,
	Z,
}

func FromString(s string) Route {
	return utils.EnumFromString(s, validRoutes, Unknown)
}

func (r Route) String() string {
	return string(r)
}

func (r *Route) Deserialize(data []byte) error {
	*r = utils.DeserializeEnum(data, FromString)
	return nil
}

func (Route) GormDataType() string {
	return gormDataTypePostgres
}

func (Route) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return gormDataTypePostgres
	default:
		return gormDataTypePostgres
	}
}

func (Route) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS ENUM (
	'1',
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'A',
	'B',
	'C',
	'D',
	'E',
	'F',
	'G',
	'J',
	'L',
	'M',
	'N',
	'Q',
	'R',
	'S',
	'SIR',
	'W',
	'Z'
);`, gormDataTypePostgres)
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of Route.String
// into a Route variable.
func (r *Route) Scan(value any) error {
	*r = utils.DbValueToEnum(value.(string), validRoutes, Unknown)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Route.String, and no error.
func (r Route) Value() (driver.Value, error) {
	return utils.EnumToDbValue(r), nil
}
