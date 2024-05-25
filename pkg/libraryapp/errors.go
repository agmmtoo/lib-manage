package libraryapp

type CoreError struct {
	Code   string
	Reason string
	Err    error
}

func (e *CoreError) Unwrap() error {
	return e.Err
}

func (e *CoreError) Error() string {
	return e.Reason
}

// Helper functions to create custom errors
func NewCoreError(code, reason string, err error) *CoreError {
	return &CoreError{
		Code:   code,
		Reason: reason,
		Err:    err,
	}
}

const (
	ErrCodeDBConnection = "ERR_DB_CONNECTION"
	ErrCodeDBQuery      = "ERR_DB_QUERY"
	ErrCodeDBExec       = "ERR_DB_EXEC"
	ErrCodeDBScan       = "ERR_DB_SCAN"
	ErrCodeDBNotFound   = "ERR_DB_NOT_FOUND"
	ErrCodeDBDuplicate  = "ERR_DB_DUPLICATE"

	ErrCodeInvalidInput = "ERR_INVALID_INPUT"
	ErrCodeMissingField = "ERR_MISSING_FIELD"
	ErrCodeInvalidEmail = "ERR_INVALID_EMAIL"
	ErrCodeDuplicate    = "ERR_DUPLICATE"

	ErrCodeAuthFailed         = "ERR_AUTH_FAILED"
	ErrCodeUnauthorized       = "ERR_UNAUTHORIZED"
	ErrCodeForbidden          = "ERR_FORBIDDEN"
	ErrCodeInternal           = "ERR_INTERNAL"
	ErrCodeServiceUnavailable = "ERR_SERVICE_UNAVAILABLE"
	ErrCodeNotImplemented     = "ERR_NOT_IMPLEMENTED"
)
