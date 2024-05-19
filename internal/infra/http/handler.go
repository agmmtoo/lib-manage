package http

import (
	"encoding/json"
	"net/http"
)

type LibraryAppHandler struct {
	service Servicer
}

func NewLibraryAppHandler(service Servicer) *LibraryAppHandler {
	return &LibraryAppHandler{service: service}
}

func (h *LibraryAppHandler) Ping(w http.ResponseWriter, r *http.Request) error {
	p, err := h.service.Ping(r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, p)
}

func (h *LibraryAppHandler) ListUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := h.service.ListUsers(r.Context(), ListUserRequest{})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, users)
}

func (h *LibraryAppHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	// if err := req.Validate(); err != nil {
	// 	return InvalidRequestData(err)
	// }

	user, err := h.service.CreateUser(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, user)
}
