package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/agmmtoo/lib-manage/internal/core/book"
	"github.com/agmmtoo/lib-manage/internal/core/user"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

type Server struct {
	addr       string
	httpServer *http.Server
}

// Business logic is implemented by the core service.
type Service interface {
	ListUsers(ctx context.Context, input user.ListRequest) (*user.ListResponse, error)
	GetUserByID(ctx context.Context, id int) (*libraryapp.User, error)

	ListBooks(ctx context.Context, input book.ListRequest) (*book.ListResponse, error)
	GetBookByID(ctx context.Context, id int) (*libraryapp.Book, error)
	// GetUsersByBookName(ctx context.Context, name string) ([]*libraryapp.User, error)
}

func NewServer(
	addr string,
	service Service,
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
	service Service,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		users, err := service.ListUsers(r.Context(), user.ListRequest{})
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
