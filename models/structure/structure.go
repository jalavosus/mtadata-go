package structure

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/basiciota"
)

//go:generate stringer -type Structure -linecomment

type Structure basiciota.BasicIota

const (
	StructureAtGrade    Structure = iota // At Grade
	StructureElevated                    // Elevated
	StructureEmbankment                  // Embankment
	StructureOpenCut                     // Open Cut
	StructureSubway                      // Subway
	StructureViaduct                     // Viaduct
	UnknownStructure                     // Unknown
)

var validStructures = []Structure{
	StructureAtGrade,
	StructureElevated,
	StructureEmbankment,
	StructureOpenCut,
	StructureSubway,
	StructureViaduct,
}

func StructureFromString(s string) Structure {
	return utils.IotaFromString(s, validStructures, UnknownStructure)
}

func (s *Structure) Deserialize(data []byte) error {
	*s = utils.DeserializeIota(data, StructureFromString)
	return nil
}

func (Structure) GormDataType() string {
	return "structure"
}

func (Structure) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return UnknownStructure.GormDataType()
	default:
		return UnknownStructure.GormDataType()
	}
}

func (Structure) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS ENUM (
	'AT_GRADE',
	'ELEVATED',
	'EMBANKMENT',
	'OPEN_CUT',
	'SUBWAY',
	'VIADUCT'
);`, UnknownStructure.GormDataType())
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of BasicIota.String
// into a Division variable.
func (s *Structure) Scan(value any) error {
	*s = utils.DbValueToIota(value.(string), validStructures, UnknownStructure)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Division.String, and no error.
func (s Structure) Value() (driver.Value, error) {
	return utils.IotaToDbValue(s), nil
}
