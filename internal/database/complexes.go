package database

import (
	"context"

	"github.com/jalavosus/mtadata/internal/database/dbconn"
	"github.com/jalavosus/mtadata/models"
)

func StationComplex(ctx context.Context, complexId string) (*models.StationComplex, error) {
	var model = &models.StationComplex{
		ComplexId: complexId,
	}

	conn := dbconn.ConnectionContext(ctx)

	if err := conn.Find(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func StationComplexes(ctx context.Context, queryParams StationComplexQueryParams) (models.StationComplexes, error) {
	var res models.StationComplexes

	params, hasParams := queryParams.ToQuery()

	conn := dbconn.ConnectionContext(ctx)

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
