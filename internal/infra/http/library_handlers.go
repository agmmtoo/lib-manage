package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *LibraryAppHandler) ListLibraries(w http.ResponseWriter, r *http.Request) error {
	libraries, err := h.service.ListLibraries(r.Context(), ListLibrariesRequest{
		// Skip: skip,
		// Limit: limit,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, libraries)
}

func (h *LibraryAppHandler) GetLibraryByID(w http.ResponseWriter, r *http.Request) error {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		return NewAPIError(http.StatusNotFound, ResourceNotFound("library"))
	}

	library, err := h.service.GetLibraryByID(r.Context(), id)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, library)
}

func (h *LibraryAppHandler) CreateLibrary(w http.ResponseWriter, r *http.Request) error {
	var req CreateLibraryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	// if err := req.Validate(); err != nil {
	// 	return InvalidRequestData(err)
	// }

	library, err := h.service.CreateLibrary(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, library)
}
