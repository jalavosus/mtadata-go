package database

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"

	"github.com/jalavosus/mtadata/internal/database/dbconn"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models"
)

var stationComplexInvalidFields = []string{
	"stations",
}

func stationComplexSelectFields(ctx context.Context) (selectFields []string) {
	if gqlFields := graphql.CollectAllFields(ctx); len(gqlFields) > 0 {
		selectFields = removeInvalidFields(gqlFields, stationComplexInvalidFields)
		if utils.SliceContains(gqlFields, "stations") {
			selectFields = append(selectFields, "station_infos")
		}
	}

	return
}

func StationComplex(ctx context.Context, complexId int64) (*models.StationComplex, error) {
	var model = &models.StationComplex{
		ComplexId: complexId,
	}

	err := dbconn.Transaction(ctx, func(tx *gorm.DB) error {
		if selectFields := stationComplexSelectFields(ctx); len(selectFields) > 0 {
			tx = tx.Select(selectFields)
		}

		if err := tx.Find(model).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return model, nil
}

func StationComplexes(ctx context.Context, queryParams StationComplexQueryParams) (models.StationComplexes, error) {
	var res models.StationComplexes

	params, hasParams := queryParams.ToQuery()
	
	err := dbconn.Transaction(ctx, func(tx *gorm.DB) error {
		tx = tx.
			Model(&models.StationComplex{}).
			Order(params.OrderBy)

		if hasParams {
			tx = tx.Where(params.Query, params.Args...)
		}

		if params.Limit != 0 {
			tx = tx.Limit(params.Limit)
		}

		if selectFields := stationComplexSelectFields(ctx); len(selectFields) > 0 {
			tx = tx.Select(selectFields)
		}

		if err := tx.Find(&res).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}