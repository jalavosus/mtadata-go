package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"

	"github.com/jalavosus/mtadata/cmd/cliutil"
)

func main() {
	app := &cli.App{
		Name:    "mtadata",
		Usage:   "Do stuff",
		Authors: []*cli.Author{&cliutil.AppAuthor},
		Commands: []*cli.Command{
			&getStationsCmd,
			&getComplexesCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
