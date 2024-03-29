package models

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jalavosus/mtadata/internal/database/dbconn"
	"github.com/jalavosus/mtadata/internal/database/dialectors"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/divisions"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
	"github.com/jalavosus/mtadata/models/routes"
)

const stationComplexGormDataTypePg string = "station_complex"

type StationComplex struct {
	Borough       boroughs.Borough    `json:"borough" yaml:"borough" gorm:"type:borough;index" gqlgen:"borough"`
	ComplexId     int64               `json:"complex_id" yaml:"complex_id" gorm:"primaryKey" gqlgen:"complex_id"`
	Divisions     divisions.Divisions `json:"divisions" yaml:"divisions" gorm:"type:division[]" gqlgen:"divisions"`
	DaytimeRoutes routes.Routes       `json:"daytime_routes" yaml:"daytime_routes" gorm:"type:route[]" gqlgen:"daytime_routes"`
	StationInfos  StationInfos        `json:"station_infos" yaml:"station_infos" gorm:"type:station_info[]" gqlgen:"station_infos"`
}

func (s StationComplex) Proto() *protosv1.StationComplex {
	return &protosv1.StationComplex{
		Borough:       s.Borough.Proto(),
		ComplexId:     s.ComplexId,
		Divisions:     s.Divisions.Proto(),
		DaytimeRoutes: s.DaytimeRoutes.Proto(),
		StationInfos:  s.StationInfos.Proto(),
	}
}

var stationComplexStationFields = []string{
	"stop_name",
	"station_id",
	"gtfs_stop_id",
	"structure",
	"line",
	"daytime_routes",
	"division",
	"gtfs_location",
	"direction_labels",
}

func (s StationComplex) Stations(ctx context.Context) (Stations, error) {
	var res Stations

	var stationIds = make([]string, len(s.StationInfos))
	for i, station := range s.StationInfos {
		stationIds[i] = station.GtfsStopId
	}

	fieldSelection := stationComplexStationFields
	if gqlFields := graphql.CollectAllFields(ctx); len(gqlFields) > 0 {
		fieldSelection = gqlFields
	}

	err := dbconn.Transaction(ctx, func(tx *gorm.DB) error {
		err := tx.
			Model(&Station{}).
			Select(fieldSelection).
			Where("gtfs_stop_id IN ? AND complex_id = ?", stationIds, s.ComplexId).
			Find(&res).
			Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (StationComplex) GormDataType() string {
	return stationComplexGormDataTypePg
}

func (StationComplex) GormDBDataType(db *gorm.DB, _ *schema.Field) string {
	switch db.Dialector.Name() {
	case dialectors.Postgres:
		return stationComplexGormDataTypePg
	default:
		return stationComplexGormDataTypePg
	}
}

func (s StationComplex) PrettyPrint() {
	fmt.Println(utils.PrettyPrintStruct(
		s,
		utils.NewPrettyPrintParam("models", "StationComplex", "StationInfo", "Station", "StationInfos"),
		utils.NewPrettyPrintParam("routes", "Route", "Routes"),
		utils.NewPrettyPrintParam("divisions", "Division", "Divisions"),
		utils.NewPrettyPrintParam("boroughs", "Borough"),
	))
}

type StationComplexes []StationComplex

func (s StationComplexes) Proto() (complexes []*protosv1.StationComplex) {
	complexes = make([]*protosv1.StationComplex, len(s))

	for i := range s {
		complexes[i] = s[i].Proto()
	}

	return
}

var (
	_ ProtoMessage[protosv1.StationComplex] = (*StationComplex)(nil)
)
