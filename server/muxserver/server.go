package muxserver

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/fx"

	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/server"
)

type MuxServer struct {
	*server.Server
}

func NewServer(conf *config.AppConfig) *MuxServer {
	endpoint := server.MakeEndpointConfig(
		conf.Server.Gateway.Host,
		conf.Server.Gateway.Port,
	)

	return &MuxServer{
		Server: server.NewServer(endpoint),
	}
}

var Module = fx.Options(
	fx.Supply(runtime.NewServeMux()),
	fx.Provide(
		NewServer,
		newHttpServer,
	),
	fx.Invoke(
		setupMux,
		serveHttp,
		runGateway,
	),
)
