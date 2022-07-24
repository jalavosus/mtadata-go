package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/jalavosus/mtadata/cmd/cliutil"
	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/internal/database"
	"github.com/jalavosus/mtadata/internal/database/dbconn"
	"github.com/jalavosus/mtadata/internal/logging"
	"github.com/jalavosus/mtadata/models"
	"github.com/jalavosus/mtadata/models/divisions"
	"github.com/jalavosus/mtadata/models/routes"
	"github.com/jalavosus/mtadata/parser"

	_ "github.com/joho/godotenv/autoload"
)

const (
	stationsOutFilename  string = "stations.json"
	complexesOutFilename string = "complexes.json"
)

var (
	insertStationsFlag = cli.BoolFlag{
		Name: "insert-stations",
		// Aliases:  []string{"s"},
		Required: false,
		Value:    false,
	}
	insertComplexesFlag = cli.BoolFlag{
		Name: "insert-complexes",
		// Aliases:  []string{"c"},
		Required: false,
		Value:    false,
	}
	configFlag = cli.PathFlag{
		Name:     "config",
		Aliases:  []string{"c"},
		Required: false,
	}
)

var (
	parseStationsCmd = cli.Command{
		Name:   "parse-stations",
		Usage:  "Parse a raw StationInfos.csv file",
		Action: parseStationsCmdAction,
	}
	parseComplexesCmd = cli.Command{
		Name:   "parse-complexes",
		Usage:  "Parse station complex data",
		Action: parseComplexesCmdAction,
	}
	readStationsJsonCmd = cli.Command{
		Name:   "read-parsed",
		Usage:  "Read the output JSON created by parse-stations",
		Action: readParsedCmdAction,
	}
	insertStationsDbCmd = cli.Command{
		Name:  "insert-db",
		Usage: "Insert output from parse-stations into Postgres",
		Flags: []cli.Flag{
			&insertStationsFlag,
			&insertComplexesFlag,
			&configFlag,
		},
		Action: insertStationsDbCmdAction,
	}
)

func parseStationsCmdAction(c *cli.Context) error {
	inputFile := c.Args().First()
	fp, fpErr := filepath.Abs(inputFile)
	if fpErr != nil {
		return fpErr
	}

	parsedStations, parseErr := parser.ParseStationsCsv(fp)
	if parseErr != nil {
		return parseErr
	}

	return writeOutputJson(parsedStations, stationsOutFilename)
}

//nolint:gocognit
func parseComplexesCmdAction(c *cli.Context) error {
	parsedStations, err := readOutputJson[models.Station](stationsOutFilename)
	if err != nil {
		return err
	}

	complexesMap := make(map[int64]*models.StationComplex)
	for _, station := range parsedStations {
		cmplx, ok := complexesMap[station.ComplexId]

		//nolint:nestif
		if ok {
			var (
				hasStation  bool
				hasDivision bool
				newRoute    bool
			)

			for _, stationInfo := range cmplx.StationInfos {
				if stationInfo.GtfsStopId == station.GtfsStopId {
					hasStation = true
				}
			}

			for _, division := range cmplx.Divisions {
				if station.Division == division {
					hasDivision = true
				}
			}

			if !hasStation {
				cmplx.StationInfos = append(cmplx.StationInfos, models.StationInfo{
					StationId:  station.StationId,
					GtfsStopId: station.GtfsStopId,
				})

				sort.Slice(cmplx.StationInfos, func(i, j int) bool {
					return cmplx.StationInfos[i].GtfsStopId < cmplx.StationInfos[j].GtfsStopId
				})
			}

			if !hasDivision {
				cmplx.Divisions = append(cmplx.Divisions, station.Division)
				sort.Slice(cmplx.Divisions, func(i, j int) bool {
					return cmplx.Divisions[i] < cmplx.Divisions[j]
				})
			}

			routesMap := make(map[routes.Route]bool)
			for _, route := range cmplx.DaytimeRoutes {
				routesMap[route] = true
			}

			for _, route := range station.DaytimeRoutes {
				if _, ok = routesMap[route]; !ok {
					routesMap[route] = true
					newRoute = true
				}
			}

			if newRoute {
				var newRoutes = make(routes.Routes, len(routesMap))

				var i = 0
				for route := range routesMap {
					newRoutes[i] = route
					i++
				}

				cmplx.DaytimeRoutes = newRoutes
				sort.Slice(cmplx.DaytimeRoutes, func(i, j int) bool {
					return cmplx.DaytimeRoutes[i] < cmplx.DaytimeRoutes[j]
				})
			}

			complexesMap[station.ComplexId] = cmplx
		} else {
			cmplx := &models.StationComplex{
				ComplexId:     station.ComplexId,
				Borough:       station.Borough,
				Divisions:     []divisions.Division{station.Division},
				DaytimeRoutes: station.DaytimeRoutes,
				StationInfos:  []models.StationInfo{{StationId: station.StationId, GtfsStopId: station.GtfsStopId}},
			}

			complexesMap[station.ComplexId] = cmplx
		}
	}

	var (
		complexes = make(models.StationComplexes, len(complexesMap))
		i         = 0
	)

	for _, cmplx := range complexesMap {
		complexes[i] = *cmplx
		i++
	}

	sort.Slice(complexes, func(i, j int) bool {
		return complexes[i].ComplexId < complexes[j].ComplexId
	})

	return writeOutputJson(complexes, complexesOutFilename)
}

func readParsedCmdAction(c *cli.Context) error {
	parsed, err := readOutputJson[models.Station](stationsOutFilename)
	if err != nil {
		return err
	}

	for _, p := range parsed {
		p.PrettyPrint()
	}

	return nil
}

func checkConstraintErr(e error) bool {
	return strings.Contains(e.Error(), "duplicate key value violates unique constraint")
}

func insert[T models.Station | models.StationComplex](conn *gorm.DB, wg *sync.WaitGroup, ch chan<- error, p []T) {
	defer wg.Done()
	err := conn.Create(p).Error

	if err != nil {
		if checkConstraintErr(err) {
			return
		}

		ch <- err
	}
}

func insertData(lc fx.Lifecycle, c *cli.Context) {
	lc.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		var (
			parsedStations  []models.Station
			parsedComplexes []models.StationComplex
			parseErr        error
		)

		insertStations := insertStationsFlag.Get(c)
		insertComplexes := insertComplexesFlag.Get(c)

		if insertStations {
			parsedStations, parseErr = readOutputJson[models.Station](stationsOutFilename)
			if parseErr != nil {
				return parseErr
			}
		}

		if insertComplexes {
			parsedComplexes, parseErr = readOutputJson[models.StationComplex](complexesOutFilename)
			if parseErr != nil {
				return parseErr
			}
		}

		// ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
		// defer cancel()

		conn := dbconn.ConnectionContext(ctx)

		errCh := make(chan error)
		var wg sync.WaitGroup

		if insertStations {
			wg.Add(1)
			go insert(conn, &wg, errCh, parsedStations)
		}

		if insertComplexes {
			wg.Add(1)
			go insert(conn, &wg, errCh, parsedComplexes)
		}

		go func() {
			defer close(errCh)
			wg.Wait()
		}()

		return <-errCh
	}})
}

func insertStationsDbCmdAction(c *cli.Context) error {
	app := fx.New(
		fx.Supply(configFlag.Get(c), c),
		fx.Provide(logging.NewLogger),
		logging.WithLogger,
		fx.Module("config", config.Module),
		fx.Module("database", database.Module),
		fx.Invoke(insertData),
		fx.StartTimeout(5*time.Minute),
	)

	if err := app.Start(c.Context); err != nil {
		return err
	}

	for {
		select {
		case <-app.Done():
			return app.Err()
		}
	}
}

func main() {
	app := &cli.App{
		Name:    "parse",
		Usage:   "Parse MTA data CSVs into JSON",
		Authors: []*cli.Author{&cliutil.AppAuthor},
		Commands: []*cli.Command{
			&parseStationsCmd,
			&readStationsJsonCmd,
			&insertStationsDbCmd,
			&parseComplexesCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
