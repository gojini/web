package web

import (
	"context"
	"net/http"
)

type Config struct {
	Address *Address `json:"address"`
	TLS     *TLS     `json:"tls,omitempty"`
}

type Server struct {
	httpServer *http.Server
	config     *Config
}

func NewServer(config *Config, handler http.Handler) *Server {
	server := &Server{
		config: config,
		httpServer: &http.Server{
			Addr:    config.Address.String(),
			Handler: handler,
		},
	}

	return server
}

func (s *Server) Start(ctx context.Context) error {
	listener, err := s.config.TLS.Listener(s.config.Address)
	if err != nil {
		return err
	}

	return s.httpServer.Serve(listener) //nolint:wrapcheck
}

func (s *Server) Stop(ctx context.Context, e error) error {
	return s.httpServer.Shutdown(ctx) //nolint:wrapcheck
}
