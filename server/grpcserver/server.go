package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
	"github.com/jalavosus/mtadata/server"
)

const (
	DefaultServerHost string = ""
	DefaultServerPort int    = 50051
)

type Server struct {
	protosv1.UnimplementedMtaDataServiceServer
	*server.Server
	grpcServer *grpc.Server
}

func NewServer() *Server {
	s := &Server{
		Server: server.NewServer(server.MakeEndpointConfig(DefaultServerHost, DefaultServerPort)),
	}

	var serverOpts []grpc.ServerOption

	s.grpcServer = grpc.NewServer(serverOpts...)

	protosv1.RegisterMtaDataServiceServer(s.grpcServer, s)

	return s
}

func (s *Server) Addr() string {
	return s.Endpoint().Addr()
}

func (s *Server) Start(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func(s *Server, ch chan<- error) {
		ch <- s.Serve()
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

func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", s.Addr())
	if err != nil {
		return errors.WithMessage(err, "error creating net.Listener instance")
	}

	s.SetStarted()
	log.Printf("MtaDataService grpc server running on %[1]s\n", s.Addr())

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
