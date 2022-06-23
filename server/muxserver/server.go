package muxserver

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/jalavosus/mtadata/server"
)

type MuxServer struct {
	*server.Server
	logger *zap.Logger
}

func NewServer(params server.NewServerParams) *MuxServer {
	endpoint := server.MakeEndpoint(params.AppConfig.Server.Gateway)

	return &MuxServer{
		Server: server.NewServer(endpoint),
		logger: params.Logger,
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
