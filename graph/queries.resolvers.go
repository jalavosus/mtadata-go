package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/jalavosus/mtadata/graph/generated"
	"github.com/jalavosus/mtadata/internal/database"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structures"
)

// Boroughs is the resolver for the boroughs field.
func (r *queryResolver) Boroughs(ctx context.Context) ([]boroughs.Borough, error) {
	return boroughs.AllBoroughs, nil
}

// Divisions is the resolver for the divisions field.
func (r *queryResolver) Divisions(ctx context.Context) ([]divisions.Division, error) {
	return divisions.AllDivisions, nil
}

// Routes is the resolver for the routes field.
func (r *queryResolver) Routes(ctx context.Context) ([]routes.Route, error) {
	return routes.AllRoutes, nil
}

// Structures is the resolver for the structures field.
func (r *queryResolver) Structures(ctx context.Context) ([]structures.Structure, error) {
	return structures.AllStructures, nil
}

// Station is the resolver for the station field.
func (r *queryResolver) Station(ctx context.Context, stationID *int, gtfsStopID *string) (*models.Station, error) {
	if stationID == nil && gtfsStopID == nil {
		return nil, ErrNoStationIdentifier
	}

	var (
		queryBy database.StationQueryBy
		queryId string
	)

	switch {
	case stationID != nil:
		queryBy = database.StationQueryByStationId
		queryId = strconv.Itoa(*stationID)
	case gtfsStopID != nil:
		queryBy = database.StationQueryByGtfsId
		queryId = *gtfsStopID
	}

	return database.Station(ctx, queryId, queryBy)
}

// Stations is the resolver for the stations field.
func (r *queryResolver) Stations(ctx context.Context, params *database.StationQueryParams) ([]models.Station, error) {
	if params == nil {
		params = new(database.StationQueryParams)
	}

	if params.OrderBy != nil {
		ob := database.OrderByFromString(string(*params.OrderBy))
		params.OrderBy = &ob
	}

	return database.Stations(ctx, *params)
}

// StationComplex is the resolver for the station_complex field.
func (r *queryResolver) StationComplex(ctx context.Context, complexID int) (*models.StationComplex, error) {
	return database.StationComplex(ctx, int64(complexID))
}

// StationComplexes is the resolver for the station_complexes field.
func (r *queryResolver) StationComplexes(ctx context.Context, params *database.StationComplexQueryParams) ([]models.StationComplex, error) {
	if params == nil {
		params = new(database.StationComplexQueryParams)
	}

	if params.OrderBy != nil {
		ob := database.OrderByFromString(string(*params.OrderBy))
		params.OrderBy = &ob
	}

	return database.StationComplexes(ctx, *params)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
