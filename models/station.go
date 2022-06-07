package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/borough"
	"github.com/jalavosus/mtadata/models/division"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structure"
)

type Station struct {
	DirectionLabels DirectionLabels     `json:"direction_labels" yaml:"direction_labels" gorm:"type:direction_labels"`
	GtfsStopId      string              `json:"gtfs_stop_id" yaml:"gtfs_stop_id" gorm:"primaryKey"`
	StopName        string              `json:"stop_name" yaml:"stop_name" gorm:"type:text"`
	Line            string              `json:"line" yaml:"line" gorm:"type:text"`
	Division        division.Division   `json:"division" yaml:"division" gorm:"index;type:division"`
	Borough         borough.Borough     `json:"borough" yaml:"borough" gorm:"index;type:borough"`
	Structure       structure.Structure `json:"structure" yaml:"structure" gorm:"type:structure"`
	DaytimeRoutes   routes.Routes       `json:"daytime_routes" yaml:"daytime_routes" gorm:"type:route[]"`
	GtfsLocation    GtfsLocation        `json:"gtfs_location" yaml:"gtfs_location" gorm:"type:gtfs_location"`
	StationId       int                 `json:"station_id" yaml:"station_id" gorm:"type:smallint;index"`
	ComplexId       int                 `json:"complex_id" yaml:"complex_id" gorm:"type:smallint"`
}

func (Station) GormDataType() string {
	return "station"
}

func (Station) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return Station{}.GormDataType()
	default:
		return Station{}.GormDataType()
	}
}

func (s Station) PrettyPrint() {
	fmt.Println(utils.PrettyPrintStruct(
		s,
		"models",
		"Station", "GtfsLocation", "DirectionLabels",
	))
}

type Stations []Station

func (Stations) GormDataType() string {
	return "station[]"
}

func (Stations) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return Stations{}.GormDataType()
	default:
		return Stations{}.GormDataType()
	}
}

func (s *Stations) Scan(value any) error {
	var stations Stations

	switch val := value.(type) {
	case Stations:
		*s = val
	case []byte:
		if err := json.Unmarshal(val, &stations); err != nil {
			return err
		}

		*s = stations
	case string:
		var ga = new(pq.StringArray)

		if err := ga.Scan(val); err != nil {
			return err
		}

		var stationMaps []map[string]any

		for _, g := range *ga {
			var station map[string]any
			if err := json.Unmarshal([]byte(g), &station); err != nil {
				return err
			}

			stationMaps = append(stationMaps, station)
		}

		for _, station := range stationMaps {
			directionLabels := station["direction_labels"].(map[string]any)
			gtfsLocation := station["gtfs_location"].(map[string]any)
			daytimeRoutes := station["daytime_routes"].([]any)

			s := Station{
				StationId:  int(station["station_id"].(float64)),
				GtfsStopId: station["gtfs_stop_id"].(string),
				StopName:   station["stop_name"].(string),
				Line:       station["line"].(string),
				Division:   division.FromString(station["division"].(string)),
				Structure:  structure.FromString(station["structure"].(string)),
				DirectionLabels: DirectionLabels{
					North: directionLabels["north"].(string),
					South: directionLabels["south"].(string),
				},
				GtfsLocation: GtfsLocation{
					Latitude:  gtfsLocation["latitude"].(float64),
					Longitude: gtfsLocation["longitude"].(float64),
				},
			}

			for _, d := range daytimeRoutes {
				s.DaytimeRoutes = append(s.DaytimeRoutes, routes.FromString(d.(string)))
			}

			stations = append(stations, s)
		}

		*s = stations
	default:
		fmt.Printf("%T\n", value)
	}

	return nil
}

func (s Stations) Value() (driver.Value, error) {
	return json.Marshal(s)
}
