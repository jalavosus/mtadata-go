package boroughs

import (
	"database/sql/driver"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/enums"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
)

//go:generate stringer -type Borough -linecomment

type Borough enums.GenericEnum[protosv1.Borough]

const (
	Unknown      Borough = iota // Unknown
	Manhattan                   // Manhattan
	Brooklyn                    // Brooklyn
	Bronx                       // Bronx
	Queens                      // Queens
	StatenIsland                // Staten Island
)

const (
	gormDataTypePg string = "borough"
)

var AllBoroughs = []Borough{
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
	return utils.EnumFromString(s, AllBoroughs, Unknown)
}

func FromProto(val protosv1.Borough) (b Borough) {
	switch val {
	case protosv1.Borough_MANHATTAN:
		b = Manhattan
	case protosv1.Borough_BROOKLYN:
		b = Brooklyn
	case protosv1.Borough_BRONX:
		b = Bronx
	case protosv1.Borough_QUEENS:
		b = Queens
	case protosv1.Borough_STATEN_ISLAND:
		b = StatenIsland
	default:
		b = Unknown
	}

	return
}

func (b Borough) Proto() protosv1.Borough {
	return protosv1.Borough(b)
}

func (Borough) GormDataType() string {
	return gormDataTypePg
}

func (Borough) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return gormDataTypePg
	default:
		return gormDataTypePg
	}
}

func (Borough) CreateDbType() string {
	return utils.MakeCreateEnumTypeCommand(AllBoroughs, gormDataTypePg)
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of Borough.String
// into a Borough variable.
func (b *Borough) Scan(value any) error {
	*b = utils.DbValueToEnum(value.(string), AllBoroughs, Unknown)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Borough.String, and no error.
func (b Borough) Value() (driver.Value, error) {
	return utils.EnumToDbValue(b), nil
}

// MarshalJSON implements json.Marshaler.
// Returns the JSON-encoded value of Borough.String.
func (b Borough) MarshalJSON() ([]byte, error) {
	return utils.SerializeEnum(b, utils.SerializeJson)
}

func (b *Borough) UnmarshalJSON(data []byte) error {
	return utils.DeserializeEnum(data, b, utils.SerializeJson, FromString)
}

func (b Borough) MarshalYAML() ([]byte, error) {
	return utils.SerializeEnum(b, utils.SerializeYaml)
}

func (b *Borough) UnmarshalYAML(data []byte) error {
	return utils.DeserializeEnum(data, b, utils.SerializeYaml, FromString)
}

func (*Borough) QueryClause() string {
	return "borough = ?"
}

func (b *Borough) Arg() *any {
	if b != nil {
		var a any = *b
		return &a
	}

	return nil
}

func (b *Borough) Invalid() bool {
	if b == nil {
		return true
	}

	return *b == Unknown
}

var (
	_ enums.Enum[protosv1.Borough] = (*Borough)(nil)
)
