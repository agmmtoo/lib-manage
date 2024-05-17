package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

type UserService interface {
	GetAllUsers() ([]*libraryapp.User, error)
}

func NewServer(
	addr string,
	userService UserService,
) (*Server, error) {
	s := &Server{
		addr: addr,
	}

	mux := http.NewServeMux()

	s.httpServer = &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	s.registerRoutes(userService)

	return s, nil
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown(context.TODO())
}

func (s *Server) registerRoutes(
	userService UserService,
) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the users to the response.
		fmt.Fprint(w, users)
	}))

	s.httpServer.Handler = mux
}
