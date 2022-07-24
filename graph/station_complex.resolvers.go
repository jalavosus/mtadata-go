package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jalavosus/mtadata/graph/generated"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
)

// Divisions is the resolver for the divisions field.
func (r *stationComplexResolver) Divisions(ctx context.Context, obj *models.StationComplex) ([]divisions.Division, error) {
	return obj.Divisions, nil
}

// DaytimeRoutes is the resolver for the daytime_routes field.
func (r *stationComplexResolver) DaytimeRoutes(ctx context.Context, obj *models.StationComplex) ([]routes.Route, error) {
	return obj.DaytimeRoutes, nil
}

// StationInfos is the resolver for the station_infos field.
func (r *stationComplexResolver) StationInfos(ctx context.Context, obj *models.StationComplex) ([]models.StationInfo, error) {
	return obj.StationInfos, nil
}

// Stations is the resolver for the stations field.
func (r *stationComplexResolver) Stations(ctx context.Context, obj *models.StationComplex) ([]models.Station, error) {
	return obj.Stations(ctx)
}

// StationComplex returns generated.StationComplexResolver implementation.
func (r *Resolver) StationComplex() generated.StationComplexResolver {
	return &stationComplexResolver{r}
}

type stationComplexResolver struct{ *Resolver }
