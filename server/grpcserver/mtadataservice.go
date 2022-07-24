package grpcserver

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jalavosus/mtadata/gtfsapi"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/apimethods"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"

	"github.com/jalavosus/mtadata/internal/database"
	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structures"
)

func (s *Server) GetStation(ctx context.Context, req *protosv1.StationRequest) (res *protosv1.StationResult, _ error) {
	res = new(protosv1.StationResult)

	var (
		queryBy database.StationQueryBy
		queryId string
	)

	var (
		stationId = req.GetStationId()
		gtfsId    = req.GetGtfsStopId()
	)

	switch {
	case stationId != "":
		queryBy = database.StationQueryByStationId
		queryId = stationId
	case gtfsId != "":
		queryBy = database.StationQueryByGtfsId
		queryId = gtfsId
	case stationId == "" && gtfsId == "":
		res.Error = models.MissingParametersError(
			apimethods.GetStation,
			"station_id", "gtfs_stop_id",
		).Proto()

		return
	}

	data, err := database.Station(ctx, queryId, queryBy)
	if err != nil {
		res.Error = dbError(err, "Station", queryId, apimethods.GetStation).Proto()
		return
	}

	res.Station = data.Proto()

	return
}

func (s *Server) GetStations(ctx context.Context, req *protosv1.StationsQuery) (res *protosv1.StationsResult, _ error) {
	res = new(protosv1.StationsResult)

	queryParams := new(database.StationQueryParams).FromProto(req)

	stations, err := database.Stations(ctx, queryParams)
	if err != nil {
		apiErr := models.InternalError(err, apimethods.GetStations)
		res.Error = apiErr.Proto()

		return
	}

	res.Stations = stations.Proto()

	return
}

func (s *Server) GetUpcomingTrains(ctx context.Context, req *protosv1.UpcomingTrainsRequest) (res *protosv1.UpcomingTrainsResult, _ error) {
	res = new(protosv1.UpcomingTrainsResult)

	var (
		stationId  = req.GetStationId()
		gtfsStopId = req.GetGtfsStopId()
		queryId    string
		queryBy    database.StationQueryBy
	)

	switch {
	case gtfsStopId != "":
		queryId = gtfsStopId
		queryBy = database.StationQueryByGtfsId
	case gtfsStopId == "" && stationId != "":
		queryId = stationId
		queryBy = database.StationQueryByStationId
	case gtfsStopId == "" && stationId == "":
		res.Error = models.MissingParametersError(
			apimethods.GetUpcomingTrains,
			"station_id", "gtfs_stop_id",
		).Proto()

		return
	}

	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if station, err := database.Station(dbCtx, queryId, queryBy); err != nil {
		res.Error = dbError(
			err,
			"Station", queryId,
			apimethods.GetUpcomingTrains,
		).Proto()

		return
	} else if gtfsStopId == "" {
		gtfsStopId = station.GtfsStopId
	}

	api := gtfsapi.NewMtaApi()

	apiCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	trainUpdates, err := api.StationStatus(apiCtx, gtfsStopId)
	if err != nil {
		res.Error = models.
			InternalError(err, apimethods.GetUpcomingTrains).
			Proto()

		return
	}

	res.UpcomingTrains = trainUpdates.Proto()

	return
}

func (s *Server) GetStationComplex(ctx context.Context, req *protosv1.StationComplexRequest) (res *protosv1.StationComplexResult, _ error) {
	res = new(protosv1.StationComplexResult)

	dbRes, err := database.StationComplex(ctx, req.GetComplexId())
	if err != nil {
		apiErr := dbError(err, "StationComplex", req.GetComplexId(), apimethods.GetStationComplex)
		s.logger.Error("error fetching complex", zap.Error(apiErr))
		res.Error = apiErr.Proto()
		return
	}

	data := dbRes.Proto()

	if req.GetVerbose() {
		var stations models.Stations

		stations, err = dbRes.Stations(ctx)
		if err != nil {
			s.logger.Error("error fetching verbose station data", zap.Error(err))
			return res, nil
		}

		data.StationInfos = nil
		data.Stations = &protosv1.Stations{
			Stations: stations.Proto(),
		}
	}

	res.StationComplex = data

	return
}

func (s *Server) GetStationComplexes(ctx context.Context, req *protosv1.StationComplexesQuery) (res *protosv1.StationComplexesResult, _ error) {
	res = new(protosv1.StationComplexesResult)

	queryParams := new(database.StationComplexQueryParams).FromProto(req)

	stationComplexes, err := database.StationComplexes(ctx, queryParams)
	if err != nil {
		res.Error = models.InternalError(err, apimethods.GetStationComplexes).Proto()
		return
	}

	var complexes = make([]*protosv1.StationComplex, len(stationComplexes))

	verbose := req.GetVerbose()

	for i, cmplx := range stationComplexes {
		data := cmplx.Proto()

		if verbose {
			var stations models.Stations

			stations, err = cmplx.Stations(ctx)
			if err != nil {
				s.logger.Error(
					"error fetching verbose station info for complex",
					zap.String("complex_id", cmplx.ComplexId),
					zap.Error(err),
				)
			} else {
				data.StationInfos = nil
				data.Stations = &protosv1.Stations{
					Stations: stations.Proto(),
				}
			}
		}

		complexes[i] = data
	}

	res.StationComplexes = complexes

	return
}

func (s *Server) GetAllRoutes(_ context.Context, _ *emptypb.Empty) (res *protosv1.AllRoutes, _ error) {
	res = new(protosv1.AllRoutes)

	res.Routes = make([]protosv1.Route, len(routes.AllRoutes))
	for i, route := range routes.AllRoutes {
		res.Routes[i] = route.Proto()
	}

	return
}

func (s *Server) GetAllBoroughs(_ context.Context, _ *emptypb.Empty) (res *protosv1.AllBoroughs, _ error) {
	res = new(protosv1.AllBoroughs)

	res.Boroughs = make([]protosv1.Borough, len(boroughs.AllBoroughs))
	for i, borough := range boroughs.AllBoroughs {
		res.Boroughs[i] = borough.Proto()
	}

	return
}

func (s *Server) GetAllDivisions(_ context.Context, _ *emptypb.Empty) (res *protosv1.AllDivisions, _ error) {
	res = new(protosv1.AllDivisions)

	res.Divisions = make([]protosv1.Division, len(divisions.AllDivisions))
	for i, division := range divisions.AllDivisions {
		res.Divisions[i] = division.Proto()
	}

	return
}

func (s *Server) GetAllStructures(_ context.Context, _ *emptypb.Empty) (res *protosv1.AllStructures, _ error) {
	res = new(protosv1.AllStructures)

	res.Structures = make([]protosv1.Structure, len(structures.AllStructures))
	for i, structure := range structures.AllStructures {
		res.Structures[i] = structure.Proto()
	}

	return
}
