package models

import (
	"github.com/jalavosus/mtadata/models/division"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structure"
)

func stringFromAny(val any) string {
	return val.(string)
}

func floatFromAny(val any) float64 {
	return val.(float64)
}

func mapFromAny(val any) map[string]any {
	return val.(map[string]any)
}

func mapSliceWithCast[T, U any](data []T, mapFn func(val T) U) (res []U) {
	res = make([]U, len(data))

	for i := range data {
		res[i] = mapFn(data[i])
	}

	return
}

func sliceFromAny[T any](vals []any, mapFn func(val any) T) (res []T) {
	res = make([]T, len(vals))

	for i := range vals {
		res[i] = mapFn(vals[i])
	}

	return
}

func stringSliceFromAny[T any](vals []any, mapFn func(val string) T) (res []T) {
	res = make([]T, len(vals))

	for i := range vals {
		res[i] = mapFn(stringFromAny(vals[i]))
	}

	return
}

func stationFromMap(m map[string]any) Station {
	directionLabels := mapFromAny(m["direction_labels"])
	gtfsLocation := mapFromAny(m["gtfs_location"])
	daytimeRoutes := m["daytime_routes"].([]any)

	return Station{
		StationId:  int(floatFromAny(m["station_id"])),
		GtfsStopId: stringFromAny(m["gtfs_stop_id"]),
		StopName:   stringFromAny(m["stop_name"]),
		Line:       stringFromAny(m["line"]),
		Division:   division.FromString(stringFromAny(m["division"])),
		Structure:  structure.FromString(stringFromAny(m["structure"])),
		DirectionLabels: DirectionLabels{
			North: stringFromAny(directionLabels["north"]),
			South: stringFromAny(directionLabels["south"]),
		},
		GtfsLocation: GtfsLocation{
			Latitude:  floatFromAny(gtfsLocation["latitude"]),
			Longitude: floatFromAny(gtfsLocation["longitude"]),
		},
		DaytimeRoutes: stringSliceFromAny(daytimeRoutes, routes.FromString),
	}
}
