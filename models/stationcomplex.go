package models

import (
	"github.com/jalavosus/mtadata/models/borough"
	"github.com/jalavosus/mtadata/models/routes"
)

type StationComplex struct {
	DaytimeRoutes routes.Routes   `json:"daytime_routes" yaml:"daytime_routes" gorm:"type:route[]"`
	Stations      Stations        `json:"stations" yaml:"stations" gorm:"type:station[]"`
	ComplexId     int             `json:"complex_id" yaml:"complex_id"`
	Borough       borough.Borough `json:"borough" yaml:"borough" gorm:"type:borough"`
}
