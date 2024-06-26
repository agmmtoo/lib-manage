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
	mux.HandleFunc("GET /api/v1/users/{id}", MakeHandler(handler.GetUserByID))

	mux.HandleFunc("GET /api/v1/books", MakeHandler(handler.ListBooks))
	mux.HandleFunc("GET /api/v1/books/{id}", MakeHandler(handler.GetBookByID))
	mux.HandleFunc("POST /api/v1/books", MakeHandler(handler.CreateBook))

	mux.HandleFunc("GET /api/v1/staffs", MakeHandler(handler.ListStaffs))
	mux.HandleFunc("GET /api/v1/staffs/{id}", MakeHandler(handler.GetStaffByID))
	mux.HandleFunc("POST /api/v1/staffs", MakeHandler(handler.CreateStaff))

	mux.HandleFunc("GET /api/v1/loans", MakeHandler(handler.ListLoans))
	mux.HandleFunc("GET /api/v1/loans/{id}", MakeHandler(handler.GetLoanByID))
	mux.HandleFunc("POST /api/v1/loans", MakeHandler(handler.CreateLoan))

	mux.HandleFunc("GET /api/v1/libraries", MakeHandler(handler.ListLibraries))
	mux.HandleFunc("GET /api/v1/libraries/{id}", MakeHandler(handler.GetLibraryByID))
	mux.HandleFunc("POST /api/v1/libraries", MakeHandler(handler.CreateLibrary))
	mux.HandleFunc("POST /api/v1/libraries/{id}/staffs", MakeHandler(handler.AssignLibraryStaff))
	mux.HandleFunc("POST /api/v1/libraries/{id}/books", MakeHandler(handler.RegisterLibraryBook))
	mux.HandleFunc("POST /api/v1/libraries/{id}/books/import", MakeHandler(handler.RegisterLibraryBookBatch))
	mux.HandleFunc("GET /api/v1/libraries/{id}/settings", MakeHandler(handler.ListLibrarySettings))
	mux.HandleFunc("PATCH /api/v1/libraries/{id}/settings", MakeHandler(handler.UpdateLibrarySettings))

	mux.HandleFunc("GET /api/v1/memberships", MakeHandler(handler.ListMemberships))

	mux.HandleFunc("GET /api/v1/public", MakeHandler(handler.GetStats))

	return mux
}
