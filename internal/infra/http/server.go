package http

import (
	"context"
	"encoding/json"
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

func (s *Server) registerRoutes(
	service Servicer,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		users, err := service.ListUsers(r.Context(), ListUserRequest{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the users to the response.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	return mux
}
