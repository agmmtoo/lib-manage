package libraryapp

type CoreError struct {
	Code   int
	Reason string
}

func (e CoreError) Error() string {
	return e.Reason
}

const (
	ENOTFOUND        = 404
	EUNAUTHORIZED    = 401
	EUNAUTHENTICATED = 403
	ECONFLICT        = 409
	EINTERNAL        = 500
)

var (
	ErrNotFound     = CoreError{Code: 404, Reason: "not found"}
	ErrInternal     = CoreError{Code: 500, Reason: "internal error"}
	ErrUnauthorized = CoreError{Code: 401, Reason: "unauthorized"}
	ErrConflict     = CoreError{Code: 409, Reason: "conflict"}
)
