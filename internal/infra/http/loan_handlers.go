package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *LibraryAppHandler) ListLoans(w http.ResponseWriter, r *http.Request) error {
	loans, err := h.service.ListLoans(r.Context(), ListLoansRequest{})
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
