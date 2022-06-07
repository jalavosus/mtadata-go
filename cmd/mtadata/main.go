package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/jalavosus/mtadata/cmd/cliutil"
	"github.com/jalavosus/mtadata/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

var (
	getComplexCmd = cli.Command{
		Name:   "get-complex",
		Usage:  "Stuff",
		Action: getComplexCmdAction,
	}
)

func getComplexCmdAction(c *cli.Context) error {
	complexId := 611

	ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
	defer cancel()

	res, err := database.StationsByComplexId(ctx, complexId)
	if err != nil {
		return err
	}

	for _, r := range res {
		r.PrettyPrint()
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:    "mtadata",
		Usage:   "Do stuff",
		Authors: []*cli.Author{&cliutil.AppAuthor},
		Commands: []*cli.Command{
			&getComplexCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
