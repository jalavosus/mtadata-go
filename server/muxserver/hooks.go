package muxserver

import (
	"context"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/jalavosus/mtadata/internal/serverauth"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
	"github.com/jalavosus/mtadata/server/grpcserver"
	"github.com/jalavosus/mtadata/server/grpcserver/compressor"
)

func newHttpServer(s *MuxServer, serveMux *runtime.ServeMux) *http.Server {
	return &http.Server{
		Addr:    s.Endpoint().Addr(),
		Handler: serveMux,
	}
}

func serveHttp(lc fx.Lifecycle, httpServer *http.Server, errCh chan error) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func(s *http.Server, ch chan<- error) {
				ch <- s.ListenAndServe()
			}(httpServer, errCh)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})
}

func runGateway(lc fx.Lifecycle, s *MuxServer, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			s.SetStarted()
			logger.Info("starting mux server", zap.String("address", s.Endpoint().Addr()))

			return nil
		},
		OnStop: func(_ context.Context) error {
			s.SetStopped()

			return nil
		},
	})
}

type SetupMuxParams struct {
	fx.In

	Server     *grpcserver.Server
	Mux        *runtime.ServeMux
	ServerAuth *serverauth.ServerAuth
}

func setupMux(lc fx.Lifecycle, logger *zap.Logger, params SetupMuxParams) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(params.ServerAuth.TransportCredentials(true)),
		grpc.WithDefaultCallOptions(
			grpc.UseCompressor(compressor.Name),
		),
	}

	startupTimeout := 20 * time.Second

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			startupCtx, cancel := context.WithTimeout(ctx, startupTimeout)
			defer cancel()

			logger.Info("waiting for grpc server to start")
		WaitStart:
			for {
				select {
				case <-startupCtx.Done():
					return startupCtx.Err()
				default:
					if params.Server.Started() {
						break WaitStart
					}
				}
			}

			err := protosv1.RegisterMtaDataServiceHandlerFromEndpoint(ctx, params.Mux, params.Server.Addr(), opts)
			if err != nil {
				return err
			}

			return nil
		},
	})
}
