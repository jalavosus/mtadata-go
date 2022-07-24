package database

import (
	"context"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

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
		sid, _ := strconv.Atoi(id)
		model.StationId = int64(sid)
	case StationQueryByGtfsId:
		model.GtfsStopId = id
	}

	err := dbconn.Transaction(ctx, func(tx *gorm.DB) error {
		if gqlFields := graphql.CollectAllFields(ctx); len(gqlFields) > 0 {
			tx = tx.Select(gqlFields)
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

func Stations(ctx context.Context, queryParams StationQueryParams) (models.Stations, error) {
	var res models.Stations

	params, hasParams := queryParams.ToQuery()

	err := dbconn.Transaction(ctx, func(tx *gorm.DB) error {
		tx = tx.
			Model(&models.Station{})

		if hasParams {
			tx = tx.Where(params.Query, params.Args...)
		}

		tx = tx.Order(clause.OrderByColumn{Column: clause.Column{Name: string(params.OrderBy)}, Desc: false})

		if gqlFields := graphql.CollectAllFields(ctx); len(gqlFields) > 0 {
			tx = tx.Select(gqlFields)
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
