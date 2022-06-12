package models

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
)

const (
	gtfsLocationGormDataTypePg string = "gtfs_location"
)

type GtfsLocation struct {
	Latitude  float64 `json:"latitude" yaml:"latitude"`
	Longitude float64 `json:"longitude" yaml:"longitude"`
}

func NewGtfsLocation(lat, long float64) GtfsLocation {
	return GtfsLocation{
		Latitude:  lat,
		Longitude: long,
	}
}

func GtfsLocationFromString(lat, long string) GtfsLocation {
	return NewGtfsLocation(utils.ParseFloat64(lat), utils.ParseFloat64(long))
}

func (g GtfsLocation) Proto() *protosv1.GtfsLocation {
	return &protosv1.GtfsLocation{
		Latitude:  g.Latitude,
		Longitude: g.Longitude,
	}
}

func (GtfsLocation) GormDataType() string {
	return gtfsLocationGormDataTypePg
}

func (GtfsLocation) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return gtfsLocationGormDataTypePg
	default:
		return gtfsLocationGormDataTypePg
	}
}

func (GtfsLocation) CreateDbType() string {
	return fmt.Sprintf(`CREATE TYPE public.%[1]s AS (
	latitude DOUBLE PRECISION,
	longitude DOUBLE PRECISION
);`, gtfsLocationGormDataTypePg)
}

func (g *GtfsLocation) Scan(value any) error {
	val := value.(string)
	val = utils.TrimParens(val)

	split := utils.TrimWhitespaceSlice(strings.Split(val, ","))

	g.Latitude = utils.ParseFloat64(split[0])
	g.Longitude = utils.ParseFloat64(split[1])

	return nil
}

func (g GtfsLocation) Value() (driver.Value, error) {
	return fmt.Sprintf("(%[1]f, %[2]f)", g.Latitude, g.Longitude), nil
}

var (
	_ ProtoMessage[protosv1.GtfsLocation] = (*GtfsLocation)(nil)
)
