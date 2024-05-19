package http

import (
	"net/http"
)

func (s *Server) registerRoutes(
	service Servicer,
) *http.ServeMux {
	mux := http.NewServeMux()

	handler := NewLibraryAppHandler(service)

	mux.HandleFunc("GET /api/v1/ping", MakeHandler(handler.Ping))

	mux.HandleFunc("GET /api/v1/users", MakeHandler(handler.ListUsers))
	mux.HandleFunc("POST /api/v1/users", MakeHandler(handler.CreateUser))

	return mux
}
