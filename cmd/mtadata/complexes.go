package main

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/jalavosus/mtadata/internal/database"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models"
)

var (
	getComplexesCmd = cli.Command{
		Name:  "complexes",
		Usage: "List station complexes with optional filtering",
		Flags: []cli.Flag{
			&routeFlag,
			&boroughFlag,
			&complexIdFlag,
			&withStationsFlag,
		},
		Action: getComplexesCmdAction,
	}
)

func getComplexesCmdAction(c *cli.Context) error {
	route, routeExists, routeValid := routeFromFlag(c)
	if !routeValid && routeExists {
		return errors.Errorf(invalidRouteMsg, string(route))
	}

	borough, boroughExists, boroughValid := boroughFromFlag(c)
	if !boroughValid && boroughExists {
		return errors.Errorf(invalidBoroughMsg, string(borough))
	}

	complexId, complexIdExists, complexIdValid := complexIdFromFlag(c)
	if !complexIdValid && complexIdExists {
		return errors.New("wtf?")
	}

	queryParams := database.StationComplexQueryParams{
		ComplexId: utils.ToPointer(complexId),
		BaseQueryParams: database.BaseQueryParams{
			Route:   utils.ToPointer(route),
			Borough: utils.ToPointer(borough),
		},
	}

	if limit := limitFlag.Get(c); limit != 0 {
		queryParams.Limit = utils.ToPointer(limit)
	}

	ctx, cancel := context.WithTimeout(c.Context, 10*time.Second)
	defer cancel()

	res, err := database.StationComplexes(ctx, queryParams)
	if err != nil {
		return err
	}

	for _, sc := range res {
		if err = outputComplex(c, sc); err != nil {
			return err
		}
	}

	return nil
}

func outputComplex(c *cli.Context, cmplx models.StationComplex) error {
	getStations := func(ctx context.Context) (models.Stations, error) {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		return cmplx.Stations(ctx)
	}

	var (
		stations    models.Stations
		stationsErr error
	)

	withStations := withStationsFlag.Get(c)

	if withStations {
		stations, stationsErr = getStations(c.Context)
		if stationsErr != nil {
			return stationsErr
		}
	}

	cmplx.PrettyPrint()
	if withStations {
		stations.PrettyPrint()
	}

	return nil
}
