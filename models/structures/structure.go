package structures

import (
	"database/sql/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/enums"
)

type Structure enums.StringEnum

const (
	AtGrade    = Structure("At Grade")
	Elevated   = Structure("Elevated")
	Embankment = Structure("Embankment")
	OpenCut    = Structure("Open Cut")
	Subway     = Structure("Subway")
	Viaduct    = Structure("Viaduct")
	Unknown    = Structure("Unknown")
)

const (
	gormDataTypePg string = "structure"
)

var AllStructures = []Structure{
	AtGrade,
	Elevated,
	Embankment,
	OpenCut,
	Subway,
	Viaduct,
}

func FromString(s string) Structure {
	return utils.EnumFromString(s, AllStructures, Unknown)
}

func (s Structure) String() string {
	return string(s)
}

func (s *Structure) Deserialize(data []byte) error {
	*s = utils.DeserializeEnum(data, FromString)
	return nil
}

func (Structure) GormDataType() string {
	return gormDataTypePg
}

func (Structure) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return gormDataTypePg
	default:
		return gormDataTypePg
	}
}

func (Structure) CreateDbType() string {
	return utils.MakeCreateEnumTypeCommand(AllStructures, gormDataTypePg)
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of BasicIota.String
// into a Division variable.
func (s *Structure) Scan(value any) error {
	*s = utils.DbValueToEnum(value.(string), AllStructures, Unknown)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Division.String, and no error.
func (s Structure) Value() (driver.Value, error) {
	return utils.EnumToDbValue(s), nil
}

func (*Structure) QueryClause() string {
	return "structure = ?"
}

func (s *Structure) Arg() *any {
	if s != nil {
		var a any = *s
		return &a
	}

	return nil
}

func (s *Structure) Invalid() bool {
	if s == nil {
		return true
	}

	return *s == Unknown
}
