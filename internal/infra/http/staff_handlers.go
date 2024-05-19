package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *LibraryAppHandler) ListStaffs(w http.ResponseWriter, r *http.Request) error {
	staffs, err := h.service.ListStaffs(r.Context(), ListStaffsRequest{})
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
		return InvalidJSON()
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
