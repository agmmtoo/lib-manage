package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *LibraryAppHandler) ListBooks(w http.ResponseWriter, r *http.Request) error {
	books, err := h.service.ListBooks(r.Context(), ListBooksRequest{})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, books)
}

func (h *LibraryAppHandler) GetBookByID(w http.ResponseWriter, r *http.Request) error {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		return NewAPIError(http.StatusNotFound, ResourceNotFound("book"))
	}

	book, err := h.service.GetBookByID(r.Context(), id)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, book)
}

func (h *LibraryAppHandler) CreateBook(w http.ResponseWriter, r *http.Request) error {
	var req CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
	}
	defer r.Body.Close()

	// if err := req.Validate(); err != nil {
	// 	return InvalidRequestData(err)
	// }

	book, err := h.service.CreateBook(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, book)
}
