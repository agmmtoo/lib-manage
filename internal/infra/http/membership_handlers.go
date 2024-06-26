package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/agmmtoo/lib-manage/pkg/libraryapp/config"
)

func (h *LibraryAppHandler) ListMemberships(w http.ResponseWriter, r *http.Request) error {

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

	var libIDs []int
	if qIDs := r.URL.Query().Get("library_ids"); qIDs != "" {
		for _, id := range strings.Split(qIDs, ",") {
			i, err := strconv.Atoi(id)
			if err != nil {
				return InvalidRequestData(map[string]string{"library_ids": "invalid"})
			}
			libIDs = append(libIDs, i)
		}
	}

	var name = r.URL.Query().Get("name")

	var dd *int
	if p := r.URL.Query().Get("duration_days"); p != "" {
		d, err := strconv.Atoi(p)
		if err != nil {
			return InvalidRequestData(map[string]string{"duration_days": "invalid"})
		}
		dd = &d
	}

	var all *int
	if p := r.URL.Query().Get("active_loan_limit"); p != "" {
		d, err := strconv.Atoi(p)
		if err != nil {
			return InvalidRequestData(map[string]string{"active_loan_limit": "invalid"})
		}
		all = &d
	}

	var fpd *int
	if p := r.URL.Query().Get("fine_per_day"); p != "" {
		d, err := strconv.Atoi(p)
		if err != nil {
			return InvalidRequestData(map[string]string{"fine_per_day": "invalid"})
		}
		fpd = &d
	}

	memberships, err := h.service.ListMemberships(r.Context(), ListMembershipsRequest{
		IDs:             ids,
		LibraryIDs:      libIDs,
		Name:            name,
		DurationDays:    dd,
		ActiveLoanLimit: all,
		FinePerDay:      fpd,
		Skip:            skip,
		Limit:           limit,
	})
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, memberships)
}
