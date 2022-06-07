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

type DirectionLabels struct {
	North string `json:"north" yaml:"north"`
	South string `json:"south" yaml:"south"`
}

func (DirectionLabels) GormDataType() string {
	return "direction_labels"
}

func (DirectionLabels) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return DirectionLabels{}.GormDataType()
	default:
		return DirectionLabels{}.GormDataType()
	}
}

func (DirectionLabels) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS (
	north TEXT,
	south TEXT
);`, DirectionLabels{}.GormDataType())
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
