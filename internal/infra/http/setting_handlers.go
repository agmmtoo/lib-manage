package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (h *LibraryAppHandler) ListLibrarySettings(w http.ResponseWriter, r *http.Request) error {
	pathID := r.PathValue("id")
	libraryID, err := strconv.Atoi(pathID)
	if err != nil {
		return NewAPIError(http.StatusNotFound, ResourceNotFound("library"))
	}

	qLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(qLimit)
	if err != nil {
		limit = 20
	}

	qSkip := r.URL.Query().Get("skip")
	skip, err := strconv.Atoi(qSkip)
	if err != nil {
		skip = 0
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

	var key = r.URL.Query().Get("key")

	res, err := h.service.ListLibrarySettings(r.Context(), ListSettingsRequest{
		IDs:        ids,
		LibraryIDs: []int{libraryID},
		Key:        key,
		Limit:      limit,
		Skip:       skip,
	})

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, res)
}

func (h *LibraryAppHandler) UpdateLibrarySettings(w http.ResponseWriter, r *http.Request) error {
	pathID := r.PathValue("id")
	_, err := strconv.Atoi(pathID)
	if err != nil {
		return NewAPIError(http.StatusNotFound, ResourceNotFound("library"))
	}

	var req UpdateSettingsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
	}
	defer r.Body.Close()

	res, err := h.service.UpdateLibrarySettings(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, res)

}
