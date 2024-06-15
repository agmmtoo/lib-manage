package http

import (
	"net/http"
)

func (h *LibraryAppHandler) GetStats(w http.ResponseWriter, r *http.Request) error {

	res, err := h.service.GetStats(r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, res)

}
