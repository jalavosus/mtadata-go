package graph

import (
	"github.com/pkg/errors"
)

var (
	ErrNoStationIdentifier = errors.New("must provide either station_id or gtfs_stop_id")
)
