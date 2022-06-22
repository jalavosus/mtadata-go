package models

import (
	"time"

	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
	"github.com/jalavosus/mtadata/models/routes"
)

type TrainTimeUpdate struct {
	GtfsStopId    string       `json:"gtfs_stop_id"`
	Station       Station      `json:"station"`
	Route         routes.Route `json:"route"`
	Direction     string       `json:"direction"`
	DepartureTime time.Time    `json:"departure_time"`
}

func (t TrainTimeUpdate) Proto() *protosv1.TrainTimeUpdate {
	return &protosv1.TrainTimeUpdate{
		GtfsStopId:    t.GtfsStopId,
		Station:       t.Station.Proto(),
		Route:         t.Route.Proto(),
		Direction:     t.Direction,
		DepartureTime: t.DepartureTime.Unix(),
	}
}

type TrainTimeUpdates []TrainTimeUpdate

func (t TrainTimeUpdates) Proto() (updates []*protosv1.TrainTimeUpdate) {
	updates = make([]*protosv1.TrainTimeUpdate, len(t))
	for i := range t {
		updates[i] = t[i].Proto()
	}

	return
}

var (
	_ ProtoMessage[protosv1.TrainTimeUpdate] = (*TrainTimeUpdate)(nil)
)
