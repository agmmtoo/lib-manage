package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/agmmtoo/lib-manage/internal/infra/config"
)

func (h *LibraryAppHandler) ListUsers(w http.ResponseWriter, r *http.Request) error {

	qLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(qLimit)
	if err != nil {
		limit = config.API_DEFAULT_LIMIT
	}

	qSkip := r.URL.Query().Get("skip")
	skip, err := strconv.Atoi(qSkip)
	if err != nil {
		skip = config.API_DEFAULT_SKIP
	}

	var ids []int
	if qIDs := r.URL.Query().Get("ids"); qIDs != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"ids": "invalid"})
			}
			ids = append(ids, i)
		}
	}

	var username = r.URL.Query().Get("username")

	users, err := h.service.ListUsers(r.Context(), ListUsersRequest{
		IDs:      ids,
		Limit:    limit,
		Skip:     skip,
		Username: username,
	})
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
