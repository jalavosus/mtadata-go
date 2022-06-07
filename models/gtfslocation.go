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

type GtfsLocation struct {
	Latitude  float64 `json:"latitude" yaml:"latitude"`
	Longitude float64 `json:"longitude" yaml:"longitude"`
}

func (GtfsLocation) GormDataType() string {
	return "gtfs_location"
}

func (GtfsLocation) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return GtfsLocation{}.GormDataType()
	default:
		return GtfsLocation{}.GormDataType()
	}
}

func (GtfsLocation) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS (
	latitude DOUBLE PRECISION,
	longitude DOUBLE PRECISION
);`, GtfsLocation{}.GormDataType())
}

func (g *GtfsLocation) Scan(value any) error {
	val := value.(string)
	val = utils.TrimParens(val)

	split := utils.TrimWhitespaceSlice(strings.Split(val, ","))

	g.Latitude = utils.ParseFloat(split[0])
	g.Longitude = utils.ParseFloat(split[1])

	return nil
}

func (g GtfsLocation) Value() (driver.Value, error) {
	return fmt.Sprintf("(%[1]f, %[2]f)", g.Latitude, g.Longitude), nil
}
