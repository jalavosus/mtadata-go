package models

import (
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structures"
)

func castString(val any) string {
	return val.(string)
}

func castFloat(val any) float64 {
	return val.(float64)
}

func castMap(val any) map[string]any {
	return val.(map[string]any)
}

type fromStringFn[T any] func(string) T

func castFromString[T any](val any, fn fromStringFn[T]) T {
	return fn(castString(val))
}

func newCastTypeFn[T any](fn fromStringFn[T]) func(any) T {
	return func(val any) T {
		return castFromString(val, fn)
	}
}

func stationFromMap(m map[string]any) Station {
	var (
		directionLabels = castMap(m["direction_labels"])
		northLabel      = castString(directionLabels["north"])
		southLabel      = castString(directionLabels["south"])
	)

	var (
		gtfsLocation = castMap(m["gtfs_location"])
		gtfsLat      = castFloat(gtfsLocation["latitude"])
		gtfsLong     = castFloat(gtfsLocation["longitude"])
	)

	daytimeRoutes := m["daytime_routes"].([]any)

	return Station{
		ComplexId:       int(castFloat(m["complex_id"])),
		StationId:       int(castFloat(m["station_id"])),
		GtfsStopId:      castString(m["gtfs_stop_id"]),
		StopName:        castString(m["stop_name"]),
		Line:            castString(m["line"]),
		Division:        castFromString(m["division"], divisions.FromString),
		Structure:       castFromString(m["structure"], structures.FromString),
		DirectionLabels: NewDirectionLabels(northLabel, southLabel),
		GtfsLocation:    NewGtfsLocation(gtfsLat, gtfsLong),
		DaytimeRoutes:   utils.MapSlice(daytimeRoutes, newCastTypeFn(routes.FromString)),
	}
}
