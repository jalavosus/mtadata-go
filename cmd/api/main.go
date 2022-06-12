package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/jalavosus/mtadata/cmd/cliutil"
	"github.com/jalavosus/mtadata/server/grpcserver"
	"github.com/jalavosus/mtadata/server/muxserver"

	_ "github.com/joho/godotenv/autoload"
)

func portFlag(port int) cli.IntFlag {
	return cli.IntFlag{
		Name:     "port",
		Usage:    "`port` for server to listen on",
		Aliases:  []string{"p"},
		Required: false,
		Value:    port,
	}
}

var (
	hostFlag = cli.StringFlag{
		Name:     "host",
		Usage:    "`host` for server to listen on",
		Required: false,
		Value:    "localhost",
	}
	grpcPortFlag = portFlag(grpcserver.DefaultServerPort)
	// restPortFlag = portFlag(8080)
)

var (
	apiGrpcCmd = cli.Command{
		Name:  "start-grpc",
		Usage: "Start the GRPC server",
		Flags: []cli.Flag{
			&hostFlag,
			&grpcPortFlag,
		},
		Action: apiGrpcCmdAction,
	}
)

func apiGrpcCmdAction(c *cli.Context) error {
	server := muxserver.NewServer()
	return server.Start(c.Context)
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
