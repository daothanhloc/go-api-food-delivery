package common

import "net/http"

type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	RootErr    error  `json:"root_error"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorizedError(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootErr
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) error {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "DB error", "DB error", "DB_ERROR")
}

func InternalServerError(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Internal server error", "Internal server error", "INTERNAL_SERVER_ERROR")
}
