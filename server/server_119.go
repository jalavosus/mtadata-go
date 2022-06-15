//go:build go1.19

package server

import (
	"strconv"
	"sync/atomic"
)

type Server struct {
	started  *atomic.Bool
	endpoint EndpointConfig
}

func NewServer(endpoint EndpointConfig) *Server {
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

func (s *Server) Endpoint() EndpointConfig {
	return s.endpoint
}

type EndpointConfig struct {
	Host string
	Port int
}

func MakeEndpointConfig(host string, port int) EndpointConfig {
	return EndpointConfig{host, port}
}

func (c EndpointConfig) Addr() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}
