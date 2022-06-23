//go:build !go1.19

package server

import (
	"strconv"
	"sync/atomic"

	"github.com/jalavosus/mtadata/internal/utils"
)

type Server struct {
	started  *uint32
	endpoint EndpointConfig
}

func NewServer(endpoint EndpointConfig) *Server {
	return &Server{
		started:  utils.ToPointer[uint32](0),
		endpoint: endpoint,
	}
}

func (s *Server) Started() bool {
	return atomic.LoadUint32(s.started) == 1
}

func (s *Server) SetStarted() {
	atomic.CompareAndSwapUint32(s.started, 0, 1)
}

func (s *Server) Stopped() bool {
	return atomic.LoadUint32(s.started) == 0
}

func (s *Server) SetStopped() {
	atomic.CompareAndSwapUint32(s.started, 1, 0)
}

func (s *Server) Endpoint() EndpointConfig {
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
	ServerAuth *serverauth.ServerAuth
	Compressor *compressor.Compressor
}
