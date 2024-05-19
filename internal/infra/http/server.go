package http

import (
	"context"
	"net/http"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

func NewServer(
	addr string,
	service Servicer,
) (*Server, error) {
	s := &Server{
		addr: addr,
	}

	s.httpServer = &http.Server{
		Addr:    s.addr,
		Handler: s.registerRoutes(service),
	}

	return s, nil
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown(context.TODO())
}
