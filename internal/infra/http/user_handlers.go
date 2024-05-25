package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *LibraryAppHandler) ListUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := h.service.ListUsers(r.Context(), ListUsersRequest{})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, users)
}

func (h *LibraryAppHandler) GetUserByID(w http.ResponseWriter, r *http.Request) error {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		return ResourceNotFound("user")
	}

	user, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, user)
}

func (h *LibraryAppHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
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
