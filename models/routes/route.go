package routes

import (
	"database/sql/driver"
	"io"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/enums"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
)

//go:generate stringer -type Route -linecomment

type Route enums.GenericEnum[protosv1.Route]

const (
	Unknown Route = iota // Unknown
	One                  // 1
	Two                  // 2
	Three                // 3
	Four                 // 4
	Five                 // 5
	Six                  // 6
	Seven                // 7
	A                    // A
	B                    // B
	C                    // C
	D                    // D
	E                    // E
	F                    // F
	G                    // G
	J                    // J
	L                    // L
	M                    // M
	N                    // N
	Q                    // Q
	R                    // R
	S                    // S
	SIR                  // SIR
	W                    // W
	Z                    // Z
)

const (
	gormDataTypePg string = "route"
)

var AllRoutes = []Route{
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
	return utils.EnumFromString(s, AllRoutes, Unknown)
}

func (r Route) IsValid() bool {
	for i := range AllRoutes {
		if AllRoutes[i] == r {
			return true
		}
	}

	return false
}

func (r Route) Proto() protosv1.Route {
	return protosv1.Route(r)
}

func (Route) GormDataType() string {
	return gormDataTypePg
}

func (Route) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return gormDataTypePg
	default:
		return gormDataTypePg
	}
}

func (Route) CreateDbType() string {
	return utils.MakeCreateEnumTypeCommand(AllRoutes, gormDataTypePg)
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of Route.String
// into a Route variable.
func (r *Route) Scan(value any) error {
	*r = utils.DbValueToEnum(value.(string), AllRoutes, Unknown)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Route.String, and no error.
func (r Route) Value() (driver.Value, error) {
	return utils.EnumToDbValue(r), nil
}

// MarshalJSON implements json.Marshaler.
// Returns the JSON-encoded value of Borough.String.
func (r Route) MarshalJSON() ([]byte, error) {
	return utils.SerializeEnum(r, utils.SerializeJson)
}

func (r *Route) UnmarshalJSON(data []byte) error {
	return utils.DeserializeEnum(data, r, utils.SerializeJson, FromString)
}

func (r Route) MarshalYAML() ([]byte, error) {
	return utils.SerializeEnum(r, utils.SerializeYaml)
}

func (r *Route) UnmarshalYAML(data []byte) error {
	return utils.DeserializeEnum(data, r, utils.SerializeYaml, FromString)
}

func (r Route) MarshalGQL(w io.Writer) {
	utils.SerializeGQL(r.String(), w)
}

func (r *Route) UnmarshalGQL(data any) error {
	return utils.DeserializeGQL(data, r, FromString)
}

func (r *Route) QueryClause() string {
	return "? = ANY(daytime_routes)"
}

func (r *Route) Arg() *any {
	if r != nil {
		var a any = *r
		return &a
	}

	return nil
}

func (r *Route) Invalid() bool {
	if r == nil {
		return true
	}

	return *r == Unknown
}

var (
	_ enums.Enum[protosv1.Route] = (*Route)(nil)
)
