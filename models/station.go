package models

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structures"
)

const (
	stationGormDataTypePg  string = "station"
	stationsGormDataTypePg string = "station[]"
)

type Station struct {
	DirectionLabels DirectionLabels      `json:"direction_labels" yaml:"direction_labels" gorm:"type:direction_labels" pp:",omitempty"`
	GtfsStopId      string               `json:"gtfs_stop_id" yaml:"gtfs_stop_id" gorm:"primaryKey"`
	StopName        string               `json:"stop_name" yaml:"stop_name" gorm:"type:text"`
	Line            string               `json:"line" yaml:"line" gorm:"type:text"`
	Division        divisions.Division   `json:"division" yaml:"division" gorm:"index;type:division"`
	Borough         boroughs.Borough     `json:"borough" yaml:"borough" gorm:"index;type:borough" pp:",omitempty"`
	Structure       structures.Structure `json:"structure" yaml:"structure" gorm:"type:structure"`
	StationId       string               `json:"station_id" yaml:"station_id" gorm:"type:text;index"`
	ComplexId       string               `json:"complex_id" yaml:"complex_id" gorm:"type:text" pp:",omitempty"`
	DaytimeRoutes   routes.Routes        `json:"daytime_routes" yaml:"daytime_routes" gorm:"type:route[]"`
	GtfsLocation    GtfsLocation         `json:"gtfs_location" yaml:"gtfs_location" gorm:"type:gtfs_location" pp:",omitempty"`
}

func (Station) GormDataType() string {
	return stationGormDataTypePg
}

func (Station) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return stationGormDataTypePg
	default:
		return stationGormDataTypePg
	}
}

func (s Station) PrettyPrint() {
	fmt.Println(utils.PrettyPrintStruct(
		s,
		utils.NewPrettyPrintParam("models", "Station", "Stations", "GtfsLocation", "DirectionLabels"),
		utils.NewPrettyPrintParam("routes", "Route", "Routes"),
		utils.NewPrettyPrintParam("divisions", "Division", "Divisions"),
		utils.NewPrettyPrintParam("boroughs", "Borough"),
	))
}

type Stations []Station

func (Stations) GormDataType() string {
	return stationsGormDataTypePg
}

func (Stations) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return stationsGormDataTypePg
	default:
		return stationsGormDataTypePg
	}
}

func (s Stations) PrettyPrint() {
	fmt.Println(utils.PrettyPrintStruct(
		s,
		utils.NewPrettyPrintParam("models", "Station", "Stations", "GtfsLocation", "DirectionLabels"),
		utils.NewPrettyPrintParam("routes", "Route", "Routes"),
		utils.NewPrettyPrintParam("divisions", "Division", "Divisions"),
		utils.NewPrettyPrintParam("boroughs", "Borough"),
	))
}
