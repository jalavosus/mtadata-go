package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/jalavosus/mtadata/cmd/cliutil"
	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/internal/database"
	"github.com/jalavosus/mtadata/internal/logging"
	"github.com/jalavosus/mtadata/internal/serverauth"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
	"github.com/jalavosus/mtadata/server/gqlserver"
	"github.com/jalavosus/mtadata/server/grpcserver"
	"github.com/jalavosus/mtadata/server/grpcserver/compressor"
	"github.com/jalavosus/mtadata/server/muxserver"
)

var (
	configFlag = cli.PathFlag{
		Name:     "config",
		Aliases:  []string{"c"},
		Required: false,
	}
	grpcServerFlag = cli.BoolFlag{
		Name:     "nogrpc",
		Usage:    "Specifies whether to start the grpc server",
		Required: false,
		Value:    false,
	}
	gqlServerFlag = cli.BoolFlag{
		Name:     "nogql",
		Usage:    "Specifies whether to start the GraphQL server",
		Required: false,
		Value:    false,
	}
	tlsCertFlag = cli.PathFlag{
		Name:     "tls-cert",
		Aliases:  []string{"t"},
		Usage:    "`path` to tls certificate",
		Required: false,
	}
	tlsKeyFlag = cli.PathFlag{
		Name:     "tls-key",
		Aliases:  []string{"k"},
		Usage:    "`path` to tls certificate key",
		Required: false,
	}
	tlsCaCertFlag = cli.PathFlag{
		Name:     "tls-ca",
		Usage:    "`path` to tls CA certificate",
		Required: false,
	}
)

var (
	clientCmd = cli.Command{
		Name:   "client",
		Usage:  "test a client",
		Action: clientCmdAction,
	}
)

func getTlsConfig(c *cli.Context) *config.TlsConfig {
	cfg := &config.TlsConfig{
		CertPath: tlsCertFlag.Get(c),
		KeyPath:  tlsKeyFlag.Get(c),
		CaPath:   tlsCaCertFlag.Get(c),
	}

	if cfg.CertPath == "" {
		cfg.UseTls = false
	} else {
		cfg.UseTls = true
	}

	return cfg
}

func clientCmdAction(c *cli.Context) error {
	var conf *config.AppConfig

	compressor.RegisterCompressor(nil)

	if confPath := configFlag.Get(c); confPath != "" {
		appConf, err := config.ReadConfig(confPath)
		if err != nil {
			return err
		}
		conf = appConf
	}

	serverAuth, err := serverauth.NewServerAuth(getTlsConfig(c))
	if err != nil {
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.UseCompressor(compressor.Name)),
		grpc.WithTransportCredentials(serverAuth.TransportCredentials(true)),
	}

	dialCtx, cancel := context.WithTimeout(c.Context, 10*time.Second)
	defer cancel()

	addr := fmt.Sprintf("%[1]s:%[2]d", conf.Server.Grpc.GetHost(), conf.Server.Grpc.GetPort())

	conn, err := grpc.DialContext(dialCtx, addr, opts...)
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()

	client := protosv1.NewMtaDataServiceClient(conn)

	queryCtx, cancel := context.WithTimeout(c.Context, 10*time.Second)
	defer cancel()

	query := &protosv1.StationComplexRequest{
		ComplexId: 611,
	}

	res, err := client.GetStationComplex(queryCtx, query)
	if err != nil {
		return err
	}

	fmt.Println(res.GetStationComplex())
	return nil
}

func apiCmdAction(c *cli.Context) error {
	errCh := make(chan error, 1)

	fxSupply := []fx.Option{
		fx.Supply(
			errCh,
			configFlag.Get(c),
			getTlsConfig(c),
		),
	}

	fxDeps := []fx.Option{
		fx.Provide(logging.NewLogger),
		logging.WithLogger,
		fx.Module("config", config.Module),
		fx.Module("database", database.Module),
		fx.Module("compressor", compressor.Module),
		fx.Module("serverauth", serverauth.Module),
	}

	if !grpcServerFlag.Get(c) {
		fxDeps = append(fxDeps, fx.Module("grpc", grpcserver.Module))
		fxDeps = append(fxDeps, fx.Module("mux", muxserver.Module))
	}
	if !gqlServerFlag.Get(c) {
		fxDeps = append(fxDeps, fx.Module("gql", gqlserver.Module))
	}

	var (
		supplyCount = len(fxSupply)
		depsCount   = len(fxDeps)
		optsCount   = supplyCount + depsCount
	)

	var fxAppOpts = make([]fx.Option, optsCount)
	for i, opt := range fxSupply {
		fxAppOpts[i] = opt
	}

	for i, opt := range fxDeps {
		fxAppOpts[i+supplyCount] = opt
	}

	app := fx.New(fxAppOpts...)

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
		Flags: []cli.Flag{
			&configFlag,
			&grpcServerFlag,
			&gqlServerFlag,
			&tlsCertFlag,
			&tlsKeyFlag,
			&tlsCaCertFlag,
		},
		Commands: []*cli.Command{
			&clientCmd,
		},
		Action: apiCmdAction,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
