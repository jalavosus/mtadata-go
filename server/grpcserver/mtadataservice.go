package grpcserver

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jalavosus/mtadata/models"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"

	"github.com/jalavosus/mtadata/internal/database"
	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structures"
)

func (s *Server) GetStation(ctx context.Context, req *protosv1.StationRequest) (*protosv1.Station, error) {
	var (
		queryBy database.StationQueryBy
		queryId string
	)

	switch {
	case req.GetStationId() != "":
		queryBy = database.StationQueryByStationId
		queryId = req.GetStationId()
	case req.GetGtfsStopId() != "":
		queryBy = database.StationQueryByGtfsId
		queryId = req.GetGtfsStopId()
	}

	res, err := database.Station(ctx, queryId, queryBy)
	if err != nil {
		return nil, err
	}

	return res.Proto(), nil
}

func (s *Server) GetStations(ctx context.Context, req *protosv1.StationsQuery) (*protosv1.StationsResult, error) {
	queryParams := new(database.StationQueryParams).FromProto(req)

	stations, err := database.Stations(ctx, queryParams)
	if err != nil {
		return nil, err
	}

	return &protosv1.StationsResult{
		Stations: stations.Proto(),
	}, nil
}

func (s *Server) GetStationComplex(ctx context.Context, req *protosv1.StationComplexRequest) (*protosv1.StationComplex, error) {
	res, err := database.StationComplex(ctx, req.GetComplexId())
	if err != nil {
		return nil, err
	}

	var stations models.Stations

	data := res.Proto()

	if req.GetVerbose() {
		stations, err = res.Stations(ctx)
		if err != nil {
			return res.Proto(), err
		}

		data.StationInfos = nil
		data.Stations = &protosv1.Stations{
			Stations: stations.Proto(),
		}
	}

	return data, nil
}

func (s *Server) GetStationComplexes(ctx context.Context, req *protosv1.StationComplexesQuery) (*protosv1.StationComplexesResult, error) {
	queryParams := new(database.StationComplexQueryParams).FromProto(req)

	stationComplexes, err := database.StationComplexes(ctx, queryParams)
	if err != nil {
		return nil, err
	}

	var res = make([]*protosv1.StationComplex, len(stationComplexes))

	verbose := req.GetVerbose()

	for i, cmplx := range stationComplexes {
		data := cmplx.Proto()

		if verbose {
			var stations models.Stations

			stations, err = cmplx.Stations(ctx)
			if err != nil {
				err = errors.WithMessagef(err, "error fetching verbose station info for complex_id %[1]s", cmplx.ComplexId)
				log.Println(err)
			} else {
				data.StationInfos = nil
				data.Stations = &protosv1.Stations{
					Stations: stations.Proto(),
				}
			}
		}

		res[i] = data
	}

	return &protosv1.StationComplexesResult{
		StationComplexes: res,
	}, nil
}

func (s *Server) GetAllRoutes(_ context.Context, _ *emptypb.Empty) (*protosv1.AllRoutes, error) {
	data := new(protosv1.AllRoutes)

	data.Routes = make([]protosv1.Route, len(routes.AllRoutes))
	for i, route := range routes.AllRoutes {
		data.Routes[i] = route.Proto()
	}

	return data, nil
}

func (s *Server) GetAllBoroughs(_ context.Context, _ *emptypb.Empty) (*protosv1.AllBoroughs, error) {
	data := new(protosv1.AllBoroughs)

	data.Boroughs = make([]protosv1.Borough, len(boroughs.AllBoroughs))
	for i, borough := range boroughs.AllBoroughs {
		data.Boroughs[i] = borough.Proto()
	}

	return data, nil
}

func (s *Server) GetAllDivisions(_ context.Context, _ *emptypb.Empty) (*protosv1.AllDivisions, error) {
	data := new(protosv1.AllDivisions)

	data.Divisions = make([]protosv1.Division, len(divisions.AllDivisions))
	for i, division := range divisions.AllDivisions {
		data.Divisions[i] = division.Proto()
	}

	return data, nil
}

func (s *Server) GetAllStructures(_ context.Context, _ *emptypb.Empty) (*protosv1.AllStructures, error) {
	data := new(protosv1.AllStructures)

	data.Structures = make([]protosv1.Structure, len(structures.AllStructures))
	for i, structure := range structures.AllStructures {
		data.Structures[i] = structure.Proto()
	}

	return data, nil
}
