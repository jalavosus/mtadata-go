package database

import (
	"context"

	"github.com/jalavosus/mtadata/internal/database/connection"
	"github.com/jalavosus/mtadata/models"
)

func Stations(ctx context.Context, queryParams StationQueryParams) (models.Stations, error) {
	var res models.Stations

	params, hasParams := queryParams.ToQuery()

	conn := connection.ConnectionContext(ctx)

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
