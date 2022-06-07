package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/jalavosus/mtadata/cmd/cliutil"
	"github.com/jalavosus/mtadata/internal/database/connection"
	"github.com/jalavosus/mtadata/internal/utils"
	"github.com/jalavosus/mtadata/models"
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
		Subcommands: []*cli.Command{
			&readStationsJsonCmd,
			&insertStationsDbCmd,
		},
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

func readOutputJson() ([]models.Station, error) {
	var parsed []models.Station

	fp, fpErr := buildParsedFilePath(stationsOutFilename)
	if fpErr != nil {
		return nil, fpErr
	}

	f, err := utils.OpenFileRead(fp)
	if err != nil {
		return nil, err
	}

	defer func() { _ = f.Close() }()

	dataBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(dataBytes, &parsed); err != nil {
		return nil, err
	}

	return parsed, nil
}

func writeOutputJson(data any, filename string) error {
	fp, fpErr := buildParsedFilePath(filename)
	if fpErr != nil {
		return fpErr
	}

	f, err := utils.OpenFileWrite(fp)
	if err != nil {
		return err
	}

	defer func() { _ = f.Close() }()

	if err = utils.ClearFile(f); err != nil {
		return err
	}

	marshalled, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(marshalled)
	if err != nil {
		return err
	}

	return nil
}

func buildParsedFilePath(filename string) (string, error) {
	joined := filepath.Join("./", "data", "parsed", filename)
	fp, fpErr := filepath.Abs(joined)

	return fp, fpErr
}

func main() {
	app := &cli.App{
		Name:    "parse",
		Usage:   "Parse MTA data CSVs into JSON",
		Authors: []*cli.Author{&cliutil.AppAuthor},
		Commands: []*cli.Command{
			&parseStationsCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
