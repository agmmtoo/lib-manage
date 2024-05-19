package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/agmmtoo/lib-manage/pkg/libraryapp"
)

type APIError struct {
	StatusCode int `json:"status_code"`
	Message    any `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	var message any = http.StatusText(http.StatusUnprocessableEntity)
	if len(errors) > 0 {
		message = errors
	}
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid JSON request data"))
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			var apiErr APIError
			var coreErr libraryapp.CoreError
			if errors.As(err, &apiErr) {
				writeJSON(w, apiErr.StatusCode, apiErr)
			} else if errors.As(err, &coreErr) {
				e := NewAPIError(coreErr.Code, coreErr)
				writeJSON(w, e.StatusCode, e)
			} else {
				errResp := map[string]any{
					"status_code": http.StatusInternalServerError,
					"message":     err.Error(),
				}
				writeJSON(w, http.StatusInternalServerError, errResp)
			}
			// slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}
