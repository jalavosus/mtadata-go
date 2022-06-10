package boroughs

import (
	"database/sql/driver"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/enums"
)

type Borough enums.StringEnum

const (
	Manhattan    = Borough("Manhattan")
	Brooklyn     = Borough("Brooklyn")
	Bronx        = Borough("Bronx")
	Queens       = Borough("Queens")
	StatenIsland = Borough("Staten Island")
	Unknown      = Borough("")
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

func (b Borough) String() string {
	return string(b)
}

func (b *Borough) Deserialize(data []byte) error {
	*b = utils.DeserializeEnum(data, FromString)
	return nil
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
