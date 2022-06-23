package grpcserver

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/jalavosus/mtadata/internal/logging"
	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
	"github.com/jalavosus/mtadata/server"
)

var Module = fx.Options(
	fx.Provide(NewServer),
	fx.Invoke(serveGrpc),
)

type Server struct {
	protosv1.UnimplementedMtaDataServiceServer
	*server.Server
	grpcServer *grpc.Server
	logger     *zap.Logger
}

func NewServer(params server.NewServerParams) (*Server, error) {
	endpoint := server.MakeEndpoint(params.AppConfig.Server.Grpc)

	s := &Server{
		Server: server.NewServer(endpoint),
		logger: params.Logger,
	}

	var serverOpts []grpc.ServerOption

	serverOpts = append(serverOpts,
		grpc.Creds(params.ServerAuth.TransportCredentials(false)),
	)

	s.grpcServer = grpc.NewServer(serverOpts...)

	protosv1.RegisterMtaDataServiceServer(s.grpcServer, s)

	return s, nil
}

func (s *Server) Addr() string {
	return s.Endpoint().Addr()
}

func (s *Server) Start(ctx context.Context) error {
	errCh := make(chan error, 1)

	logger := logging.NewLogger()

	go func(s *Server, ch chan<- error) {
		ch <- s.Serve(logger)
	}(s, errCh)

	for {
		select {
		case <-ctx.Done():
			s.Stop(true)
			return nil
		case err := <-errCh:
			return err
		}
	}
}

func (s *Server) Serve(logger *zap.Logger) error {
	lis, err := net.Listen("tcp", s.Addr())
	if err != nil {
		return errors.WithMessage(err, "error creating net.Listener instance")
	}

	s.logger = logger
	s.SetStarted()
	logger.Info("starting grpc server", zap.String("address", s.Endpoint().Addr()))

	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop(graceful bool) {
	if graceful {
		s.grpcServer.GracefulStop()
	} else {
		s.grpcServer.Stop()
	}

	s.SetStopped()
}

func serveGrpc(lc fx.Lifecycle, s *Server, logger *zap.Logger, errCh chan error) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func(s *Server, logger *zap.Logger, ch chan<- error) {
				ch <- s.Serve(logger)
			}(s, logger, errCh)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.Stop(true)
			return nil
		},
	})
}
