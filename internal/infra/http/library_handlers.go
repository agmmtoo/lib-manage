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
		return InvalidJSON(err)
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

func (h *LibraryAppHandler) AssignLibraryStaff(w http.ResponseWriter, r *http.Request) error {
	var req AssignLibraryStaffRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
	}
	defer r.Body.Close()

	pathID := r.PathValue("id")
	libraryID, err := strconv.Atoi(pathID)
	if err != nil {
		return NewAPIError(http.StatusNotFound, ResourceNotFound("library"))
	}

	res, err := h.service.AssignLibraryStaff(r.Context(), AssignLibraryStaffRequest{
		LibraryID: libraryID,
		StaffID:   req.StaffID,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, res)
}

func (h *LibraryAppHandler) RegisterLibraryBook(w http.ResponseWriter, r *http.Request) error {
	var req RegisterLibraryBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
	}
	defer r.Body.Close()

	pathID := r.PathValue("id")
	libraryID, err := strconv.Atoi(pathID)
	if err != nil {
		return NewAPIError(http.StatusNotFound, ResourceNotFound("library"))
	}

	res, err := h.service.RegisterLibraryBook(r.Context(), RegisterLibraryBookRequest{
		LibraryID: libraryID,
		BookID:    req.BookID,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, res)
}

func (h *LibraryAppHandler) RegisterLibraryBookBatch(w http.ResponseWriter, r *http.Request) error {
	var req RegisterLibraryBookBatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
	}
	defer r.Body.Close()

	pathID := r.PathValue("id")
	libraryID, err := strconv.Atoi(pathID)
	if err != nil {
		return NewAPIError(http.StatusNotFound, ResourceNotFound("library"))
	}

	res, err := h.service.RegisterLibraryBookBatch(r.Context(), RegisterLibraryBookBatchRequest{
		LibraryID:      libraryID,
		BookIDs:        req.BookIDs,
		SkipDuplicates: req.SkipDuplicates,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, res)
}
