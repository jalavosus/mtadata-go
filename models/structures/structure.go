package structures

import (
	"database/sql/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/enums"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
)

//go:generate stringer -type Structure -linecomment

type Structure enums.GenericEnum[protosv1.Structure]

const (
	Unknown    Structure = iota // Unknown
	AtGrade                     // At Grade
	Elevated                    // Elevated
	Embankment                  // Embankment
	OpenCut                     // Open Cut
	Subway                      // Subway
	Viaduct                     // Viaduct
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

func FromProto(val protosv1.Structure) (s Structure) {
	switch val {
	case protosv1.Structure_AT_GRADE:
		s = AtGrade
	case protosv1.Structure_ELEVATED:
		s = Elevated
	case protosv1.Structure_EMBANKMENT:
		s = Embankment
	case protosv1.Structure_OPEN_CUT:
		s = OpenCut
	case protosv1.Structure_SUBWAY:
		s = Subway
	case protosv1.Structure_VIADUCT:
		s = Viaduct
	default:
		s = Unknown
	}

	return
}

func (s Structure) Proto() protosv1.Structure {
	return protosv1.Structure(s)
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

// MarshalJSON implements json.Marshaler.
// Returns the JSON-encoded value of Division.String.
func (s Structure) MarshalJSON() ([]byte, error) {
	return utils.SerializeEnum(s, utils.SerializeJson)
}

func (s *Structure) UnmarshalJSON(data []byte) error {
	return utils.DeserializeEnum(data, s, utils.SerializeJson, FromString)
}

func (s Structure) MarshalYAML() ([]byte, error) {
	return utils.SerializeEnum(s, utils.SerializeYaml)
}

func (s *Structure) UnmarshalYAML(data []byte) error {
	return utils.DeserializeEnum(data, s, utils.SerializeYaml, FromString)
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

var (
	_ enums.Enum[protosv1.Structure] = (*Structure)(nil)
)
