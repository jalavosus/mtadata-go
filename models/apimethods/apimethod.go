package apimethods

//go:generate stringer -type ApiMethod -linecomment

type ApiMethod uint

const (
	GetStation          ApiMethod = iota // GetStation
	GetStations                          // GetStations
	GetUpcomingTrains                    // GetUpcomingTrains
	GetStationComplex                    // GetStationComplex
	GetStationComplexes                  // GetStationComplexes
	GetAllRoutes                         // GetAllRoutes
	GetAllBoroughs                       // GetAllBoroughs
	GetAllDivisions                      // GetAllDivisions
	GetAllStructures                     // GetAllStructures
)
