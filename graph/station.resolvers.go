package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jalavosus/mtadata/graph/generated"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/routes"
)

// DaytimeRoutes is the resolver for the daytime_routes field.
func (r *stationResolver) DaytimeRoutes(ctx context.Context, obj *models.Station) ([]routes.Route, error) {
	panic(fmt.Errorf("not implemented"))
}

// Station returns generated.StationResolver implementation.
func (r *Resolver) Station() generated.StationResolver { return &stationResolver{r} }

type stationResolver struct{ *Resolver }
