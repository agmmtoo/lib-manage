package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
			var coreErr *libraryapp.CoreError
			if errors.As(err, &coreErr); coreErr != nil {
				statusCode := http.StatusBadRequest
				switch coreErr.Code {

				case libraryapp.ErrCodeDBNotFound:
					statusCode = http.StatusNotFound

				case libraryapp.ErrCodeDBScan,
					libraryapp.ErrCodeDBExec,
					libraryapp.ErrCodeDBDuplicate,
					libraryapp.ErrCodeDBQuery:
					statusCode = http.StatusBadRequest
				}

				errResp := map[string]any{
					"code":    coreErr.Code,
					"message": coreErr.Reason,
				}
				log.Println(errors.Unwrap(err))
				writeJSON(w, statusCode, errResp)
			} else if apiErr, ok := err.(APIError); ok {
				errResp := map[string]any{
					"code":    "ERR_API",
					"message": apiErr.Message,
				}
				log.Println(apiErr)
				writeJSON(w, apiErr.StatusCode, errResp)
			} else {
				errResp := map[string]any{
					"code":    libraryapp.ErrCodeInternal,
					"message": err.Error(),
				}
				log.Println("hi", err)
				writeJSON(w, http.StatusInternalServerError, errResp)
			}
			// slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func ResourceNotFound(resource string) APIError {
	return NewAPIError(http.StatusNotFound, fmt.Errorf("%s not found", resource))
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}
