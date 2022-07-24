package main

import (
	"github.com/urfave/cli/v2"

	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/routes"
)

var (
	routeFlag = cli.StringFlag{
		Name:     "route",
		Usage:    "Filter stations and complexes by `route`",
		Aliases:  []string{"r"},
		Required: false,
	}
	boroughFlag = cli.StringFlag{
		Name:     "borough",
		Usage:    "Filter stations and complexes by `borough`",
		Aliases:  []string{"b"},
		Required: false,
	}
	complexIdFlag = cli.Int64Flag{
		Name:     "complex-id",
		Usage:    "Filter stations and complexes by `complex_id`",
		Aliases:  []string{"c"},
		Required: false,
	}
	limitFlag = cli.IntFlag{
		Name:     "limit",
		Usage:    "Limit number of returned items",
		Aliases:  []string{"l"},
		Required: false,
	}
	withStationsFlag = cli.BoolFlag{
		Name:     "with-stations",
		Usage:    "Output verbose data for stations in a complex",
		Aliases:  []string{"s"},
		Required: false,
		Value:    false,
	}
)

func complexIdFromFlag(c *cli.Context) (complexId int64, exists, valid bool) {
	val := complexIdFlag.Get(c)
	exists = val != 0

	if exists {
		complexId = val
		valid = true
	}

	return
}

func routeFromFlag(c *cli.Context) (route routes.Route, exists, valid bool) {
	route = routes.Unknown

	val := routeFlag.Get(c)
	exists = val != ""

	if exists {
		route = routes.FromString(val)
		valid = route != routes.Unknown
	}

	return
}

func boroughFromFlag(c *cli.Context) (borough boroughs.Borough, exists, valid bool) {
	borough = boroughs.Unknown

	val := boroughFlag.Get(c)
	exists = val != ""

	if exists {
		borough = boroughs.FromString(val)
		valid = borough != boroughs.Unknown
	}

	return
}
