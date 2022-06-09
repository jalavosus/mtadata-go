package models

import (
	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/routes"
)

type StationComplex struct {
	Borough       boroughs.Borough `json:"borough" yaml:"borough" gorm:"type:borough"`
	DaytimeRoutes routes.Routes    `json:"daytime_routes" yaml:"daytime_routes" gorm:"type:route[]"`
	Stations      Stations         `json:"stations" yaml:"stations" gorm:"type:station[]"`
	ComplexId     int              `json:"complex_id" yaml:"complex_id"`
}
