package database

import (
	"context"

	"github.com/jalavosus/mtadata/internal/database/dbconn"
	"github.com/jalavosus/mtadata/models"
)

type StationQueryBy uint

const (
	StationQueryByStationId StationQueryBy = iota
	StationQueryByGtfsId
)

func Station(ctx context.Context, id string, queryBy StationQueryBy) (*models.Station, error) {
	var model = &models.Station{}

	switch queryBy {
	case StationQueryByStationId:
		model.StationId = id
	case StationQueryByGtfsId:
		model.GtfsStopId = id
	}

	conn := dbconn.ConnectionContext(ctx)

	if err := conn.Find(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func Stations(ctx context.Context, queryParams StationQueryParams) (models.Stations, error) {
	var res models.Stations

	params, hasParams := queryParams.ToQuery()

	conn := dbconn.ConnectionContext(ctx)

	tx := conn.Model(&models.Station{})
	if hasParams {
		tx = tx.Where(params.Query, params.Args...)
	}

	tx = tx.Order(params.OrderBy)

	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
