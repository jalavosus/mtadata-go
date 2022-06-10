package models

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
)

const (
	directionLabelsGormDataTypePg string = "direction_labels"
)

type DirectionLabels struct {
	North string `json:"north" yaml:"north" pp:",omitempty"`
	South string `json:"south" yaml:"south" pp:",omitempty"`
}

func NewDirectionLabels(north, south string) DirectionLabels {
	return DirectionLabels{North: north, South: south}
}

func (DirectionLabels) GormDataType() string {
	return directionLabelsGormDataTypePg
}

func (DirectionLabels) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return directionLabelsGormDataTypePg
	default:
		return directionLabelsGormDataTypePg
	}
}

func (DirectionLabels) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS (
	north TEXT,
	south TEXT
);`, directionLabelsGormDataTypePg)
}

func (d *DirectionLabels) Scan(value any) error {
	val := value.(string)
	val = utils.TrimParens(val)

	split := utils.TrimWhitespaceSlice(strings.Split(val, ","))

	d.North = strings.TrimSpace(strings.ReplaceAll(split[0], `"`, ""))
	d.South = strings.TrimSpace(strings.ReplaceAll(split[1], `"`, ""))

	return nil
}

func (d DirectionLabels) Value() (driver.Value, error) {
	return fmt.Sprintf(`("%[1]s", "%[2]s")`, d.North, d.South), nil
}
