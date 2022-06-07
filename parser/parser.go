package parser

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/borough"
	"github.com/jalavosus/mtadata/models/division"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structure"
)

func ParseStationsCsv(csvPath string) ([]models.Station, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}

	defer func() { _ = f.Close() }()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	records = records[1:]

	parsed := make([]models.Station, len(records))

	for i, record := range records {
		parsed[i] = parseCsvRecord(record)
	}

	return parsed, nil
}

func parseCsvRecord(record []string) models.Station {
	daytimeRoutesRaw := strings.Split(record[7], " ")
	daytimeRoutes := make(routes.Routes, len(daytimeRoutesRaw))

	for i, rt := range daytimeRoutesRaw {
		daytimeRoutes[i] = routes.RouteFromString(rt)
	}

	return models.Station{
		StationId:     parseInt(record[0]),
		ComplexId:     parseInt(record[1]),
		GtfsStopId:    record[2],
		Division:      division.DivisionFromString(record[3]),
		Line:          record[4],
		StopName:      record[5],
		Borough:       borough.BoroughFromMtaCsv(record[6]),
		DaytimeRoutes: daytimeRoutes,
		Structure:     structure.StructureFromString(record[8]),
		GtfsLocation: models.GtfsLocation{
			Latitude:  parseFloat(record[9]),
			Longitude: parseFloat(record[10]),
		},
		DirectionLabels: models.DirectionLabels{
			North: record[11],
			South: record[12],
		},
	}
}

func parseInt(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
