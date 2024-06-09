package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/agmmtoo/lib-manage/pkg/libraryapp/config"
)

func (h *LibraryAppHandler) ListStaffs(w http.ResponseWriter, r *http.Request) error {
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

	var userIDs []int
	if qIDs := r.URL.Query().Get("user_ids"); qIDs != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"user_ids": "invalid"})
			}
			userIDs = append(userIDs, i)
		}
	}

	var libraryIDs []int
	if lIDs := r.URL.Query().Get("library_ids"); lIDs != "" {
		for _, id := range strings.Split(lIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"library_ids": "invalid"})
			}
			libraryIDs = append(libraryIDs, i)
		}
	}

	staffs, err := h.service.ListStaffs(r.Context(), ListStaffsRequest{
		IDs:        ids,
		UserIDs:    userIDs,
		Limit:      limit,
		Skip:       skip,
		LibraryIDs: libraryIDs,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, staffs)
}

func (h *LibraryAppHandler) GetStaffByID(w http.ResponseWriter, r *http.Request) error {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		return ResourceNotFound("staff")
	}

	staff, err := h.service.GetStaffByID(r.Context(), id)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, staff)
}

func (h *LibraryAppHandler) CreateStaff(w http.ResponseWriter, r *http.Request) error {
	var req CreateStaffRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
	}
	defer r.Body.Close()

	// if err := req.Validate(); err != nil {
	// 	return InvalidRequestData(err)
	// }

	staff, err := h.service.CreateStaff(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, staff)
}
