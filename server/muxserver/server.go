package muxserver

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/server"
)

type MuxServer struct {
	*server.Server
	logger *zap.Logger
}

func NewServer(params server.NewServerParams) *MuxServer {
	endpointConf := params.AppConfig.Server.Gateway
	if endpointConf.Port == 0 {
		(&endpointConf).Port = config.DefaultPortGateway
	}

	endpoint := server.MakeEndpoint(endpointConf)

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
