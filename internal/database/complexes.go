package database

import (
	"context"

	"github.com/jalavosus/mtadata/internal/database/connection"
	"github.com/jalavosus/mtadata/models"
)

func StationComplexes(ctx context.Context, queryParams StationComplexQueryParams) (models.StationComplexes, error) {
	var res models.StationComplexes

	params, hasParams := queryParams.ToQuery()

	conn := connection.ConnectionContext(ctx)

	tx := conn.
		Model(&models.StationComplex{}).
		Order(params.OrderBy)

	if hasParams {
		tx = tx.Where(params.Query, params.Args...)
	}

	if params.Limit != 0 {
		tx = tx.Limit(params.Limit)
	}

	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
