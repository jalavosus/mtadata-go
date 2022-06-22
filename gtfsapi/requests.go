package gtfsapi

import (
	"github.com/jalavosus/mtadata/models/routes"
)

const (
	apiEndpoint string = "https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs"
)

func endpointForRoute(route routes.Route) string {
	var routeEndpoint string

	switch route {
	case routes.G:
		routeEndpoint = "-g"
	case routes.L:
		routeEndpoint = "-l"
	case routes.SIR:
		routeEndpoint = "-si"
	case routes.J, routes.Z:
		routeEndpoint = "-jz"
	case routes.A, routes.C, routes.E:
		routeEndpoint = "-ace"
	case routes.B, routes.D, routes.F, routes.M:
		routeEndpoint = "-bdfm"
	case routes.N, routes.Q, routes.R, routes.W:
		routeEndpoint = "-nqrw"
	case routes.One, routes.Two, routes.Three, routes.Four,
		routes.Five, routes.Six, routes.Seven:
		routeEndpoint = ""
	}

	return routeEndpoint
}
