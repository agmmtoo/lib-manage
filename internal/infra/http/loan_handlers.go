package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/agmmtoo/lib-manage/pkg/libraryapp/config"
)

func (h *LibraryAppHandler) ListLoans(w http.ResponseWriter, r *http.Request) error {
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

	qActive := r.URL.Query().Get("active")
	active, err := strconv.ParseBool(qActive)
	if err != nil {
		active = false
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
	if qIDs := r.URL.Query().Get("user_ids"); strings.TrimSpace(qIDs) != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"user_ids": "invalid"})
			}
			userIDs = append(userIDs, i)
		}

	}

	var bookIDS []int
	if qIDs := r.URL.Query().Get("book_ids"); strings.TrimSpace(qIDs) != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"book_ids": "invalid"})
			}
			bookIDS = append(bookIDS, i)
		}

	}

	var libIDs []int
	if qIDs := r.URL.Query().Get("library_ids"); strings.TrimSpace(qIDs) != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"library_ids": "invalid"})
			}
			libIDs = append(libIDs, i)
		}

	}

	var staffIDs []int
	if qIDs := r.URL.Query().Get("staff_ids"); strings.TrimSpace(qIDs) != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"staff_ids": "invalid"})
			}
			staffIDs = append(staffIDs, i)
		}

	}

	var includeLibraryBook, includeSubscription, includeStaff bool
	if qInclude := r.URL.Query().Get("include"); strings.TrimSpace(qInclude) != "" {
		for _, inc := range strings.Split(qInclude, ",") {
			switch inc {
			case "book":
				includeLibraryBook = true
			case "subscription":
				includeSubscription = true
			case "staff":
				includeStaff = true
			}
		}
	}

	loans, err := h.service.ListLoans(r.Context(), ListLoansRequest{
		Limit:               limit,
		Skip:                skip,
		Active:              active,
		IDs:                 ids,
		UserIDs:             userIDs,
		LibraryBookIDs:      bookIDS,
		LibraryIDs:          libIDs,
		StaffIDs:            staffIDs,
		IncludeLibraryBook:  includeLibraryBook,
		IncludeSubscription: includeSubscription,
		IncludeStaff:        includeStaff,
	})

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, loans)
}

func (h *LibraryAppHandler) GetLoanByID(w http.ResponseWriter, r *http.Request) error {
	pathID := r.PathValue("id")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		return ResourceNotFound("loan")
	}

	loan, err := h.service.GetLoanByID(r.Context(), id)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, loan)
}

func (h *LibraryAppHandler) CreateLoan(w http.ResponseWriter, r *http.Request) error {
	var req CreateLoanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON(err)
	}
	defer r.Body.Close()

	// if err := req.Validate(); err != nil {
	// 	return InvalidRequestData(err)
	// }

	loan, err := h.service.CreateLoan(r.Context(), req)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, loan)
}
