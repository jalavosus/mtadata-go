//go:build go1.19

package server

import (
	"strconv"
	"sync/atomic"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/jalavosus/mtadata/internal/config"
	"github.com/jalavosus/mtadata/internal/serverauth"
	"github.com/jalavosus/mtadata/server/grpcserver/compressor"
)

type Server struct {
	started  *atomic.Bool
	endpoint Endpoint
}

func NewServer(endpoint Endpoint) *Server {
	started := new(atomic.Bool)

	return &Server{
		started:  started,
		endpoint: endpoint,
	}
}

func (s *Server) Started() bool {
	return s.started.Load()
}

func (s *Server) SetStarted() {
	s.started.CompareAndSwap(false, true)
}

func (s *Server) Stopped() bool {
	return !s.started.Load()
}

func (s *Server) SetStopped() {
	s.started.CompareAndSwap(true, false)
}

func (s *Server) Endpoint() Endpoint {
	return s.endpoint
}

type Endpoint struct {
	Host string
	Port int
}

func MakeEndpoint(endpointConfig config.EndpointConfig) Endpoint {
	return Endpoint{
		Host: endpointConfig.GetHost(),
		Port: endpointConfig.GetPort(),
	}
}

func (e Endpoint) Addr() string {
	return e.Host + ":" + strconv.Itoa(e.Port)
}

type NewServerParams struct {
	fx.In

	Logger     *zap.Logger
	AppConfig  *config.AppConfig
	Compressor *compressor.Compressor
	ServerAuth *serverauth.ServerAuth
}
