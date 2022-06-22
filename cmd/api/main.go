package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"

	_ "github.com/joho/godotenv/autoload"

	"github.com/jalavosus/mtadata/cmd/cliutil"
	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/internal/database"
	"github.com/jalavosus/mtadata/internal/logging"
	"github.com/jalavosus/mtadata/server/grpcserver"
	"github.com/jalavosus/mtadata/server/muxserver"
)

var (
	configFlag = cli.PathFlag{
		Name:     "config",
		Aliases:  []string{"c"},
		Required: false,
	}
)

var (
	apiGrpcCmd = cli.Command{
		Name:  "start-grpc",
		Usage: "Start the GRPC server",
		Flags: []cli.Flag{
			&configFlag,
		},
		Action: apiGrpcCmdAction,
	}
)

func apiGrpcCmdAction(c *cli.Context) error {
	errCh := make(chan error, 1)

	app := fx.New(
		fx.Supply(errCh, configFlag.Get(c)),
		fx.Provide(logging.NewLogger),
		logging.WithLogger,
		fx.Module("config", config.Module),
		fx.Module("grpc", grpcserver.Module),
		fx.Module("mux", muxserver.Module),
		fx.Module("database", database.Module),
	)

	if err := app.Start(c.Context); err != nil {
		return err
	}

	for {
		select {
		case <-app.Done():
			if app.Err() != nil {
				return app.Err()
			}

			return nil
		case err := <-errCh:
			return err
		}
	}
}

func main() {
	app := &cli.App{
		Name:    "api",
		Usage:   "Start the API server(s)",
		Authors: []*cli.Author{&cliutil.AppAuthor},
		Commands: []*cli.Command{
			&apiGrpcCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
