package main

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/jalavosus/mtadata/internal/database"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/boroughs"
)

var (
	getStationsCmd = cli.Command{
		Name:  "stations",
		Usage: "List stations with optional filtering",
		Flags: []cli.Flag{
			&routeFlag,
			&boroughFlag,
		},
		Action: getStationsCmdAction,
	}
)

const (
	invalidRouteMsg string = "invalid route name %[1]s"
)

var (
	invalidBoroughMsg = "invalid borough name %[1]s; allowed boroughs: \n" +
		strings.Join(
			utils.ToStringSlice(boroughs.AllBoroughs),
			",\n",
		)
)

func getStationsCmdAction(c *cli.Context) error {
	route, routeExists, routeValid := routeFromFlag(c)
	if !routeValid && routeExists {
		return errors.Errorf(invalidRouteMsg, string(route))
	}

	borough, boroughExists, boroughValid := boroughFromFlag(c)
	if !boroughValid && boroughExists {
		return errors.Errorf(invalidBoroughMsg, string(borough))
	}

	queryParams := database.StationQueryParams{
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

	res, err := database.Stations(ctx, queryParams)
	if err != nil {
		return err
	}

	for _, station := range res {
		station.PrettyPrint()
	}

	return nil
}
