package http

import (
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
