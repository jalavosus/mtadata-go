package gtfsapi

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/jalavosus/mtadata/internal/database/dbconn"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/protos/gtfs"
	"github.com/jalavosus/mtadata/models/routes"
)

type MtaApi struct {
	httpClient *http.Client
}

func NewMtaApi() *MtaApi {
	return &MtaApi{
		httpClient: new(http.Client),
	}
}

func (m MtaApi) StationStatus(ctx context.Context, gtfsStopId string) (models.TrainTimeUpdates, error) {
	var trainStatuses models.TrainTimeUpdates

	station, err := m.linesForStation(ctx, gtfsStopId)
	if err != nil {
		return nil, err
	}

	for _, route := range station.DaytimeRoutes {
		feed, err := m.routeData(ctx, route)
		if err != nil {
			return nil, err
		}

		for _, feedItem := range feed.GetEntity() {
			tripUpdate := feedItem.GetTripUpdate()
			if tripUpdate == nil {
				continue
			}

			trip := tripUpdate.GetTrip()

			routeId := trip.GetRouteId()
			if route.String() != routeId {
				continue
			}

			// var nyctTrip = new(gtfs.NyctTripDescriptor)
			// b := trip.ProtoReflect()
			// if err = proto.Unmarshal(b, nyctTrip); err != nil {
			// 	return err
			// }

			stopTimeUpdates := tripUpdate.GetStopTimeUpdate()
			if stopTimeUpdates == nil {
				continue
			}

			for _, stu := range stopTimeUpdates {
				if strings.HasPrefix(stu.GetStopId(), gtfsStopId) {
					ttu := models.TrainTimeUpdate{
						GtfsStopId:    gtfsStopId,
						Station:       *station,
						Route:         routes.FromString(routeId),
						DepartureTime: time.Unix(stu.GetDeparture().GetTime(), 0),
					}

					trainStatuses = append(trainStatuses, ttu)
				}
			}
		}

		// fmt.Println(trainStatuses)
	}

	return trainStatuses, nil
}

func (m MtaApi) routeData(ctx context.Context, route routes.Route) (*gtfs.FeedMessage, error) {
	routeEndpoint := endpointForRoute(route)

	endpoint := apiEndpoint + routeEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	apiKey := mtaApiKey()
	// apiKeyB64 := base64.StdEncoding.EncodeToString([]byte(apiKey))

	req.Header.Set(apiKeyHeader, apiKey)

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data = new(gtfs.FeedMessage)

	if err = proto.Unmarshal(respBody, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (m MtaApi) linesForStation(ctx context.Context, gtfsStopId string) (*models.Station, error) {
	station := models.Station{GtfsStopId: gtfsStopId}

	err := dbconn.
		ConnectionContext(ctx).
		First(&station).
		Error

	if err != nil {
		return nil, err
	}

	return &station, nil
}
