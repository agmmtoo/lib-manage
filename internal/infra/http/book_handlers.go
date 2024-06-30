package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/agmmtoo/lib-manage/pkg/libraryapp/config"
)

func (h *LibraryAppHandler) ListLibraryBooks(w http.ResponseWriter, r *http.Request) error {
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

	var title = r.URL.Query().Get("title")
	var author = r.URL.Query().Get("author")

	var libraryIDs []int
	if qIDs := r.URL.Query().Get("library_ids"); qIDs != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"library_ids": "invalid"})
			}
			libraryIDs = append(libraryIDs, i)
		}
	}

	books, err := h.service.ListLibraryBooks(r.Context(), ListLibraryBooksRequest{
		IDs:        ids,
		LibraryIDs: libraryIDs,
		Title:      title,
		Author:     author,
		Limit:      limit,
		Skip:       skip,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, books)
}

func (h *LibraryAppHandler) GetLibraryBookByID(w http.ResponseWriter, r *http.Request) error {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		return NewAPIError(http.StatusNotFound, ResourceNotFound("book"))
	}

	book, err := h.service.GetLibraryBookByID(r.Context(), id)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, book)
}

func (h *LibraryAppHandler) CreateLibraryBook(w http.ResponseWriter, r *http.Request) error {
	var req CreateLibraryBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
	}
	defer r.Body.Close()

	// if err := req.Validate(); err != nil {
	// 	return InvalidRequestData(err)
	// }

	book, err := h.service.CreateLibraryBook(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, book)
}
