package parser

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/boroughs"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/models/structures"
)

func ParseStationsCsv(csvPath string) (stations []models.Station, err error) {
	csvLines, err := readCsv(csvPath, true)
	if err != nil {
		return
	}

	stations = make([]models.Station, len(csvLines))

	for i, record := range csvLines {
		stations[i] = parseStationFromCsv(record)
	}

	return
}

func parseStationFromCsv(record []string) models.Station {
	return models.Station{
		StationId:       utils.ParseInt(record[0]),
		ComplexId:       utils.ParseInt(record[1]),
		GtfsStopId:      record[2],
		Division:        divisions.FromString(record[3]),
		Line:            record[4],
		StopName:        record[5],
		Borough:         boroughs.FromMtaCsvString(record[6]),
		DaytimeRoutes:   utils.MapSlice(strings.Split(record[7], " "), routes.FromString),
		Structure:       structures.FromString(record[8]),
		GtfsLocation:    models.GtfsLocationFromString(record[9], record[10]),
		DirectionLabels: models.NewDirectionLabels(record[11], record[12]),
	}
}

func readCsv(csvPath string, removeHeader bool) (records [][]string, err error) {
	var f *os.File

	f, err = utils.OpenFileRead(utils.AbsolutePath(csvPath))
	if err != nil {
		return
	}

	defer func() {
		_ = f.Close()
	}()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	records, err = reader.ReadAll()
	if err != nil {
		return
	}

	// remove first row from result slice
	// if removeHeader is true, as it's
	// (likely) a header row and contains nothing useful to us.
	if removeHeader {
		records = records[1:]
	}

	return
}
