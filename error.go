package barerror

import (
	"errors"
	"fmt"
	"net/http"
)

// AppError represents application specific error.
type AppError struct {
	Message string
	Code    int
	Err     error
}

// predefined error types
var (
	ErrValidation = newType("validation error", http.StatusBadRequest)
	ErrDBWrite    = newType("database error: unable to write to db", http.StatusInternalServerError)
	ErrNotFound   = newType("resource not found", http.StatusNotFound)
)

func newType(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
		Err:     errors.New(message),
	}
}

func (e *AppError) New(message string) *AppError {
	return &AppError{
		Message: fmt.Sprintf("%s: %s", e.Message, message),
		Code:    e.Code,
		Err:     fmt.Errorf("%w: %s", e.Err, message),
	}
}

// Wrap wraps err with a custom message.
func (e *AppError) Wrap(err error, msg string) *AppError {
	return &AppError{
		Message: fmt.Sprintf("%s: %s", e.Message, msg),
		Code:    e.Code,
		Err:     fmt.Errorf("%s: %s: %w", e.Message, msg, err),
	}
}

// Error implements the error interface for AppError.
func (e *AppError) Error() string {
	return fmt.Sprintf("Error: %v", e.Err)
}

// Unwrap returns the underlying error.
func (e *AppError) Unwrap() error {
	return e.Err
}
