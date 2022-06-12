package muxserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	protosv1 "github.com/jalavosus/mtadata/models/protos/v1"
	"github.com/jalavosus/mtadata/server"
	"github.com/jalavosus/mtadata/server/grpcserver"
)

const (
	DefaultHost string = ""
	DefaultPort int    = 9091
)

type MuxServer struct {
	*server.Server
}

func NewServer() *MuxServer {
	return &MuxServer{
		Server: server.NewServer(server.MakeEndpointConfig(DefaultHost, DefaultPort)),
	}
}

func (s *MuxServer) Start(ctx context.Context) error {
	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, os.Kill, os.Interrupt)

	grpcServer := grpcserver.NewServer()

	go func(s *grpcserver.Server, ch chan<- error) {
		errCh <- grpcServer.Serve()
	}(grpcServer, errCh)

	log.Println("waiting for grpc server to be started...")
	for !grpcServer.Started() {
	}

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := protosv1.RegisterMtaDataServiceHandlerFromEndpoint(ctx, mux, grpcServer.Addr(), opts)
	if err != nil {
		return err
	}

	httpServer := &http.Server{
		Addr:    s.Endpoint().Addr(),
		Handler: mux,
	}

	go func(ctx context.Context, s *http.Server, ch chan<- error) {
		ch <- s.ListenAndServe()
	}(ctx, httpServer, errCh)

	cleanup := func(err error) error {
		grpcServer.Stop(true)

		if stopErr := httpServer.Shutdown(ctx); stopErr != nil {
			return stopErr
		}

		s.SetStopped()

		return err
	}

	s.SetStarted()
	log.Printf("starting mux server on %[1]s\n", s.Endpoint().Addr())

	for {
		select {
		case <-ctx.Done():
			return cleanup(nil)
		case sig := <-sigCh:
			sigErr := errors.Errorf("received signal %[1]s from os", sig.String())
			return cleanup(sigErr)
		case serverErr := <-errCh:
			return cleanup(serverErr)
		}
	}
}
