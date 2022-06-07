package database

import (
	"context"

	"github.com/jalavosus/mtadata/internal/database/connection"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/routes"
)

const (
	AllComplexesQuery string = `SELECT DISTINCT ON (complex_id)
	complex_id,
	boroughs[1] AS borough,
	daytime_routes,
	ARRAY(SELECT * from json_array_elements(stations)) AS stations
FROM
	(
		SELECT
			complex_id,
			array_agg(DISTINCT borough ORDER BY borough) AS boroughs,
			array_agg(DISTINCT routes ORDER BY routes) AS daytime_routes,
			json_agg((
				 SELECT station FROM
					  (
							SELECT
								 station_id,
								 gtfs_stop_id,
								 stop_name,
								 daytime_routes,
								 division,
								 line,
								 direction_labels,
								 structure,
								 gtfs_location
					  ) AS station
			)) AS stations
		FROM
			(
				SELECT DISTINCT ON (gtfs_stop_id)
					complex_id,
					station_id,
					gtfs_stop_id,
					gtfs_location,
					stop_name,
					line,
					borough,
					structure,
					unnest(daytime_routes) AS routes,
					daytime_routes,
					division,
					direction_labels
				FROM stations
			) s
		GROUP BY complex_id
) a
ORDER BY complex_id;`

	ComplexByIdQuery string = `SELECT
	complex_id,
	boroughs[1] AS borough,
	daytime_routes,
	ARRAY(SELECT * from json_array_elements(stations)) AS stations
FROM
	(
		SELECT
			complex_id,
			array_agg(DISTINCT borough ORDER BY borough) AS boroughs,
			array_agg(DISTINCT routes ORDER BY routes) AS daytime_routes,
			json_agg((
				 SELECT station FROM
					  (
							SELECT
								 station_id,
								 gtfs_stop_id,
								 stop_name,
								 daytime_routes,
								 division,
								 line,
								 direction_labels,
								 structure,
								 gtfs_location
					  ) AS station
			)) AS stations
		FROM
			(
				SELECT DISTINCT ON (gtfs_stop_id)
					complex_id,
					station_id,
					gtfs_stop_id,
					gtfs_location,
					stop_name,
					line,
					borough,
					structure,
					unnest(daytime_routes) AS routes,
					daytime_routes,
					division,
					direction_labels
				FROM stations
			) s
		GROUP BY complex_id
) a
WHERE complex_id = ?`

	AllStationsByComplexQuery string = `SELECT DISTINCT ON (a.complex_id)
    a.complex_id,
    a.station_ids,
    a.gtfs_stop_ids,
    a.stop_names,
    a.daytime_routes,
    a.divisions,
    b.borough
FROM
(
    SELECT
        complex_id,
        array_agg(DISTINCT stop_name ORDER BY stop_name) AS stop_names,
        array_agg(DISTINCT routes ORDER BY routes) AS daytime_routes,
        array_agg(DISTINCT station_id ORDER BY station_id) AS station_ids,
        array_agg(DISTINCT gtfs_stop_id ORDER BY gtfs_stop_id) AS gtfs_stop_ids,
        array_agg(DISTINCT division ORDER BY division) AS divisions
    FROM
    (
        SELECT
            complex_id,
            station_id,
            gtfs_stop_id,
            stop_name,
            unnest(daytime_routes) AS routes,
            division
        FROM stations
    ) s
    GROUP BY complex_id
) a
RIGHT JOIN stations b USING (complex_id)
WHERE a.complex_id = ?
ORDER BY complex_id, gtfs_stop_id;`

	StationByComplexIdQuery string = `SELECT DISTINCT ON (a.station_id)
    a.station_id,
    a.complex_id,
    a.gtfs_stop_ids,
    a.stop_names,
    a.daytime_routes,
    a.divisions,
    b.borough
FROM
(
    SELECT
        station_id,
        max(complex_id) AS complex_id,
        array_agg(DISTINCT stop_name ORDER BY stop_name) AS stop_names,
        array_agg(DISTINCT routes ORDER BY routes) AS daytime_routes,
        array_agg(DISTINCT gtfs_stop_id ORDER BY gtfs_stop_id) AS gtfs_stop_ids,
        array_agg(DISTINCT division ORDER BY division) AS divisions
    FROM
    (
        SELECT
            complex_id,
            station_id,
            gtfs_stop_id,
            stop_name,
            unnest(daytime_routes) AS routes,
            division
        FROM stations
    ) s
    GROUP BY station_id
) a
RIGHT JOIN stations b USING (station_id)
WHERE a.complex_id = ?
ORDER BY station_id;`

	StationsByComplexIdQuery string = `SELECT *
FROM stations
WHERE complex_id = ?
ORDER BY gtfs_stop_id;`

	StationsByRouteAndComplexIdQuery string = `SELECT *
FROM stations
WHERE ? = ANY(daytime_routes)
AND complex_id = ?
ORDER BY gtfs_stop_id;`

	StationsByRouteQuery string = `SELECT *
FROM stations
WHERE ? = ANY(daytime_routes)
ORDER BY gtfs_stop_id;`
)

func StationComplexes(ctx context.Context) ([]models.StationComplex, error) {
	var res []models.StationComplex

	conn := connection.ConnectionContext(ctx)

	tx := conn.Raw(AllComplexesQuery)
	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// func StationsForComplex(ctx context.Context, complexId int) (*models.StationComplex, error) {
// 	var res *models.StationComplex
//
// 	if err := fetchFromRawQuery(ctx, AllStationsByComplexQuery, &res, complexId); err != nil {
// 		return nil, err
// 	}
//
// 	return res, nil
// }

func StationComplexById(ctx context.Context, complexId int) (*models.StationComplex, error) {
	var res models.StationComplex

	if err := fetchFromRawQuery(ctx, ComplexByIdQuery, &res, complexId); err != nil {
		return nil, err
	}

	return &res, nil
}

// func StationByComplexId(ctx context.Context, complexId int) (*models.Station, error) {
// 	var res models.Station
//
// 	if err := fetchFromRawQuery(ctx, StationByComplexIdQuery, &res, complexId); err != nil {
// 		return nil, err
// 	}
//
// 	return &res, nil
// }

func StationsByComplexId(ctx context.Context, complexId int) (models.Stations, error) {
	var res models.Stations

	if err := fetchFromRawQuery(ctx, StationsByComplexIdQuery, &res, complexId); err != nil {
		return nil, err
	}

	return res, nil
}

func StationsByRoute(ctx context.Context, route routes.Route) (models.Stations, error) {
	var res models.Stations

	if err := fetchFromRawQuery(ctx, StationsByRouteQuery, &res, route); err != nil {
		return nil, err
	}

	return res, nil
}

func StationsByRouteAndComplexId(ctx context.Context, complexId int, route routes.Route) (models.Stations, error) {
	var res models.Stations

	if err := fetchFromRawQuery(ctx, StationsByRouteAndComplexIdQuery, &res, complexId, route); err != nil {
		return nil, err
	}

	return res, nil
}

func fetchFromRawQuery(ctx context.Context, rawQuery string, out any, args ...any) error {
	conn := connection.ConnectionContext(ctx)

	err := conn.
		Raw(rawQuery, args...).
		Scan(out).
		Error

	if err != nil {
		return err
	}

	return nil
}
