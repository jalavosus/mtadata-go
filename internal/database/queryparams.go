package database

import (
	"strings"

	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"

	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structures"
)

type OrderBy string

const (
	OrderByGtfsStopId = OrderBy("gtfs_stop_id")
	OrderByStationId  = OrderBy("station_id")
	OrderByComplexId  = OrderBy("complex_id")
	OrderByUnknown    = OrderBy("")
)

type QueryParam[T any] interface {
	QueryClause() string
	Arg() *T
	Invalid() bool
}

func checkAppendParam[T comparable](param *T, clause string, queryParams []string, args []any) ([]string, []any) {
	var zeroValue T

	if p, ok := checkValid(param, zeroValue); ok {
		queryParams = append(queryParams, clause)
		args = append(args, p)
	}

	return queryParams, args
}

func checkAppendQueryParam[T any](param QueryParam[T], queryParams []string, args []any) ([]string, []any) {
	if p, ok := checkValidParam(param); ok {
		queryParams = append(queryParams, param.QueryClause())
		args = append(args, p)
	}

	return queryParams, args
}

type QueryParams struct {
	Query   string
	OrderBy OrderBy
	Args    []any
	Limit   int
}

type BaseQueryParams struct {
	Route     *routes.Route
	Borough   *boroughs.Borough
	Division  *divisions.Division
	Structure *structures.Structure
	OrderBy   *OrderBy
	Limit     *int
}

func (p BaseQueryParams) ToQuery(queryParams []string, args []any) (params QueryParams, hasParams bool) {
	queryParamSlice := []QueryParam[any]{
		p.Route,
		p.Borough,
		p.Division,
		p.Structure,
	}

	for _, qp := range queryParamSlice {
		queryParams, args = checkAppendQueryParam(qp, queryParams, args)
	}

	if limit, ok := checkValid(p.Limit, 0); ok {
		params.Limit = limit
	}

	if len(queryParams) > 0 && len(args) > 0 {
		params.Query = strings.Join(queryParams, " AND ")
		params.Args = args
		hasParams = true
	}

	return
}

func (p *BaseQueryParams) FromProto(params *protosv1.BaseQueryParams, orderByStation bool) BaseQueryParams {
	baseParams := BaseQueryParams{}

	if params == nil {
		return baseParams
	}

	if params.Borough != nil {
		baseParams.Borough = utils.ToPointer(boroughs.FromProto(params.GetBorough()))
	}

	if params.Division != nil {
		baseParams.Division = utils.ToPointer(divisions.FromProto(params.GetDivision()))
	}

	// if params.Route != nil {
	// 	p.Route = utils.ToPointer(routes.)
	// }

	if params.OrderBy != nil {
		switch params.GetOrderBy() {
		case protosv1.QueryOrderBy_GTFS_STOP_ID:
			baseParams.OrderBy = utils.ToPointer(OrderByGtfsStopId)
		case protosv1.QueryOrderBy_STATION_ID:
			baseParams.OrderBy = utils.ToPointer(OrderByStationId)
		case protosv1.QueryOrderBy_COMPLEX_ID:
			baseParams.OrderBy = utils.ToPointer(OrderByComplexId)
		default:
			if orderByStation {
				baseParams.OrderBy = utils.ToPointer(OrderByGtfsStopId)
			} else {
				baseParams.OrderBy = utils.ToPointer(OrderByComplexId)
			}
		}
	}

	if params.Limit != nil && params.GetLimit() != 0 {
		baseParams.Limit = utils.ToPointer(int(params.GetLimit()))
	}

	return baseParams
}

type StationQueryParams struct {
	BaseQueryParams
	StationId  *string
	ComplexId  *string
	GtfsStopId *string
}

func (p *StationQueryParams) FromProto(params *protosv1.StationsQuery) StationQueryParams {
	var baseParams BaseQueryParams

	if params.QueryParams != nil {
		baseParams = new(BaseQueryParams).
			FromProto(params.QueryParams.BaseParams, true)
	}

	return StationQueryParams{
		BaseQueryParams: baseParams,
		StationId:       params.QueryParams.StationId,
		ComplexId:       params.QueryParams.ComplexId,
		GtfsStopId:      params.QueryParams.GtfsStopId,
	}
}

func (p StationQueryParams) ToQuery() (params QueryParams, hasParams bool) {
	var (
		queryParams []string
		args        []any
	)

	queryParams, args = checkAppendParam(p.StationId, "station_id = ?", queryParams, args)
	queryParams, args = checkAppendParam(p.ComplexId, "complex_id = ?", queryParams, args)
	queryParams, args = checkAppendParam(p.GtfsStopId, "gtfs_stop_id = ?", queryParams, args)

	params, hasParams = p.BaseQueryParams.ToQuery(queryParams, args)
	params.OrderBy = OrderByGtfsStopId

	if orderBy, ok := checkValid(p.OrderBy, OrderByUnknown); ok {
		params.OrderBy = orderBy
	}

	return
}

type StationComplexQueryParams struct {
	BaseQueryParams
	ComplexId *string
}

func (p *StationComplexQueryParams) FromProto(params *protosv1.StationComplexesQuery) StationComplexQueryParams {
	var baseParams BaseQueryParams

	if params.QueryParams != nil {
		baseParams = new(BaseQueryParams).
			FromProto(params.QueryParams.BaseParams, false)
	}

	return StationComplexQueryParams{
		BaseQueryParams: baseParams,
		ComplexId:       params.QueryParams.ComplexId,
	}
}

func (p StationComplexQueryParams) ToQuery() (params QueryParams, hasParams bool) {
	var (
		queryParams []string
		args        []any
	)

	queryParams, args = checkAppendParam(p.ComplexId, "complex_id = ?", queryParams, args)

	params, hasParams = p.BaseQueryParams.ToQuery(queryParams, args)
	params.OrderBy = OrderByComplexId

	if orderBy, ok := checkValid(p.OrderBy, OrderByUnknown); ok {
		params.OrderBy = orderBy
	}

	return
}
