package routes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/basiciota"
)

//go:generate stringer -type Route -linecomment

type Route basiciota.BasicIota

const (
	Line1        Route = iota // 1
	Line2                     // 2
	Line3                     // 3
	Line4                     // 4
	Line5                     // 5
	Line6                     // 6
	Line7                     // 7
	LineA                     // A
	LineB                     // B
	LineC                     // C
	LineD                     // D
	LineE                     // E
	LineF                     // F
	LineG                     // G
	LineJ                     // J
	LineL                     // L
	LineM                     // M
	LineN                     // N
	LineQ                     // Q
	LineR                     // R
	LineS                     // S
	SIR                       // SIR
	LineW                     // W
	LineZ                     // Z
	UnknownRoute              // Unknown
)

var validRoutes = []Route{
	Line1,
	Line2,
	Line3,
	Line4,
	Line5,
	Line6,
	Line7,
	LineA,
	LineB,
	LineC,
	LineD,
	LineE,
	LineF,
	LineG,
	LineJ,
	LineL,
	LineM,
	LineN,
	LineQ,
	LineR,
	LineS,
	SIR,
	LineW,
	LineZ,
}

func RouteFromString(s string) Route {
	return utils.IotaFromString(s, validRoutes, UnknownRoute)
}

func (r *Route) Deserialize(data []byte) error {
	*r = utils.DeserializeIota(data, RouteFromString)
	return nil
}

func (Route) GormDataType() string {
	return "route"
}

func (Route) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return UnknownRoute.GormDataType()
	default:
		return UnknownRoute.GormDataType()
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
);`, UnknownRoute.GormDataType())
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of Route.String
// into a Route variable.
func (r *Route) Scan(value any) error {
	*r = utils.DbValueToIota(value.(string), validRoutes, UnknownRoute)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Route.String, and no error.
func (r Route) Value() (driver.Value, error) {
	return utils.IotaToDbValue(r), nil
}

type Routes []Route

func (r *Routes) UnmarshalJSON(data []byte) error {
	var routes []Route
	if err := json.Unmarshal(data, &routes); err != nil {
		return err
	}

	*r = routes

	return nil
}

func (r Routes) Value() (driver.Value, error) {
	var sa = make(pq.StringArray, len(r))

	for i, rt := range r {
		sa[i] = rt.String()
	}

	return sa.Value()
}

func (r *Routes) Scan(value any) error {
	var sa = new(pq.StringArray)

	if err := sa.Scan(value); err != nil {
		return err
	}

	routes := make(Routes, len(*sa))
	for i, rt := range *sa {
		routes[i] = utils.DbValueToIota(rt, validRoutes, UnknownRoute)
	}

	*r = routes

	return nil
}

func (Routes) GormDataType() string {
	return "route[]"
}

func (Routes) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return Routes{}.GormDataType()
	default:
		return Routes{}.GormDataType()
	}
}
