package divisions

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
)

const (
	divisionsGormDataTypePg string = "division[]"
)

type Divisions []Division

func (d *Divisions) UnmarshalJSON(data []byte) error {
	var divisions []Division
	if err := json.Unmarshal(data, &divisions); err != nil {
		return err
	}

	*d = divisions

	return nil
}

func (d Divisions) Value() (driver.Value, error) {
	var sa = make(pq.StringArray, len(d))

	for i, rt := range d {
		sa[i] = rt.String()
	}

	return sa.Value()
}

func (d *Divisions) Scan(value any) error {
	var sa = new(pq.StringArray)

	if err := sa.Scan(value); err != nil {
		return err
	}

	routes := make(Divisions, len(*sa))
	for i, rt := range *sa {
		routes[i] = utils.DbValueToEnum(rt, AllDivisions, Unknown)
	}

	*d = routes

	return nil
}

func (Divisions) GormDataType() string {
	return divisionsGormDataTypePg
}

func (Divisions) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return divisionsGormDataTypePg
	default:
		return divisionsGormDataTypePg
	}
}
