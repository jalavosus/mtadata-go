package gqlserver

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/server"
)

const (
	serverEndpoint string = "/graphql"
)

type GqlServer struct {
	*server.Server
	logger *zap.Logger
}

func NewServer(params server.NewServerParams) *GqlServer {
	endpointConf := params.AppConfig.Server.Graphql
	if endpointConf.Port == 0 {
		(&endpointConf).Port = config.DefaultPortGraphql
	}

	endpoint := server.MakeEndpoint(endpointConf)

	return &GqlServer{
		Server: server.NewServer(endpoint),
		logger: params.Logger,
	}
}

var Module = fx.Options(
	fx.Provide(
		NewServer,
		newGqlHandler,
		newRouter,
		newHttpServer,
	),
	fx.Invoke(
		runGql,
	),
)
