package divisions

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

//go:generate stringer -type Division -linecomment

type Division enums.GenericEnum[protosv1.Division]

const (
	Unknown Division = iota // Unknown
	BMT                     // BMT
	IND                     // IND
	IRT                     // IRT
	SIR                     // SIR
)

const (
	gormDataTypePg string = "division"
)

var AllDivisions = []Division{
	BMT,
	IND,
	IRT,
	SIR,
}

func FromString(s string) Division {
	return utils.EnumFromString(s, AllDivisions, Unknown)
}

func FromProto(val protosv1.Division) (d Division) {
	switch val {
	case protosv1.Division_BMT:
		d = BMT
	case protosv1.Division_IND:
		d = IND
	case protosv1.Division_IRT:
		d = IRT
	case protosv1.Division_SIR_DIVISION:
		d = SIR
	default:
		d = Unknown
	}

	return
}

func (d Division) IsValid() bool {
	for i := range AllDivisions {
		if AllDivisions[i] == d {
			return true
		}
	}

	return false
}

func (d Division) Proto() protosv1.Division {
	return protosv1.Division(d)
}

func (Division) GormDataType() string {
	return gormDataTypePg
}

func (Division) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return gormDataTypePg
	default:
		return gormDataTypePg
	}
}

func (Division) CreateDbType() string {
	return utils.MakeCreateEnumTypeCommand(AllDivisions, gormDataTypePg)
}

// Scan implements sql.Scanner.
// Sets the driver.Value represenation of BasicIota.String
// into a Division variable.
func (d *Division) Scan(value any) error {
	*d = utils.DbValueToEnum(value.(string), AllDivisions, Unknown)
	return nil
}

// Value implements driver.Valuer.
// Returns the result of Division.String, and no error.
func (d Division) Value() (driver.Value, error) {
	return utils.EnumToDbValue(d), nil
}

// MarshalJSON implements json.Marshaler.
// Returns the JSON-encoded value of Division.String.
func (d Division) MarshalJSON() ([]byte, error) {
	return utils.SerializeEnum(d, utils.SerializeJson)
}

func (d *Division) UnmarshalJSON(data []byte) error {
	return utils.DeserializeEnum(data, d, utils.SerializeJson, FromString)
}

func (d Division) MarshalYAML() ([]byte, error) {
	return utils.SerializeEnum(d, utils.SerializeYaml)
}

func (d *Division) UnmarshalYAML(data []byte) error {
	return utils.DeserializeEnum(data, d, utils.SerializeYaml, FromString)
}

func (d Division) MarshalGQL(w io.Writer) {
	utils.SerializeGQL(d.String(), w)
}

func (d *Division) UnmarshalGQL(data any) error {
	return utils.DeserializeGQL(data, d, FromString)
}

func (d *Division) QueryClause() string {
	return "division = ?"
}

func (d *Division) Arg() *any {
	if d != nil {
		var a any = *d
		return &a
	}

	return nil
}

func (d *Division) Invalid() bool {
	if d == nil {
		return true
	}

	return *d == Unknown
}

var (
	_ enums.Enum[protosv1.Division] = (*Division)(nil)
)
