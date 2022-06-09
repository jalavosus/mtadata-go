package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/jalavosus/mtadata/cmd/cliutil"
	"github.com/jalavosus/mtadata/internal/database/connection"
	"github.com/jalavosus/mtadata/parser"

	_ "github.com/joho/godotenv/autoload"
)

const (
	stationsOutFilename string = "stations.json"
)

var (
	parseStationsCmd = cli.Command{
		Name:   "parse-stations",
		Usage:  "Parse a raw Stations.csv file",
		Action: parseStationsCmdAction,
	}
	readStationsJsonCmd = cli.Command{
		Name:   "read-parsed",
		Usage:  "Read the output JSON created by parse-stations",
		Action: readParsedCmdAction,
	}
	insertStationsDbCmd = cli.Command{
		Name:   "insert-db",
		Usage:  "Insert output from parse-stations into Postgres",
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

func readParsedCmdAction(c *cli.Context) error {
	parsed, err := readOutputJson()
	if err != nil {
		return err
	}

	for _, p := range parsed {
		p.PrettyPrint()
	}

	return nil
}

func insertStationsDbCmdAction(c *cli.Context) error {
	parsed, err := readOutputJson()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
	defer cancel()

	conn := connection.ConnectionContext(ctx)

	for _, p := range parsed {
		if err = conn.Create(&p).Error; err != nil {
			if strings.Contains(err.Error(), `duplicate key value violates unique constraint "stations_pkey"`) {
				continue
			}

			return err
		}
	}

	return nil
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
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
