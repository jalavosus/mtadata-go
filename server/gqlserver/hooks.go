package gqlserver

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/jalavosus/mtadata/graph/generated"

	"github.com/jalavosus/mtadata/graph"
)

func newGqlHandler(s *GqlServer) *handler.Server {
	conf := generated.Config{Resolvers: &graph.Resolver{
		Logger: s.logger,
	}}

	schema := generated.NewExecutableSchema(conf)

	gqlHandler := handler.NewDefaultServer(schema)

	return gqlHandler
}

func newRouter(gqlHandler *handler.Server) *mux.Router {
	r := mux.NewRouter()
	playgroundHandler := playground.Handler("GraphQL Playground", serverEndpoint)
	r.Handle("/", playgroundHandler)
	r.Handle(serverEndpoint, gqlHandler)

	return r
}

func newHttpServer(s *GqlServer, router *mux.Router) *http.Server {
	return &http.Server{
		Addr:    s.Endpoint().Addr(),
		Handler: router,
	}
}

func runGql(lc fx.Lifecycle, s *GqlServer, logger *zap.Logger, httpServer *http.Server, errCh chan error) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func(s *http.Server, ch chan<- error) {
				ch <- s.ListenAndServe()
			}(httpServer, errCh)

			s.SetStarted()
			logger.Info("starting gql server", zap.String("address", s.Endpoint().Addr()))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.SetStopped()

			return httpServer.Shutdown(ctx)
		},
	})
}
