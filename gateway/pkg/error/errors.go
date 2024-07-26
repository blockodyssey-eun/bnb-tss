package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

var (
	ErrInvalidJSON    = AppError{Status: http.StatusBadRequest, Message: "Invalid JSON payload"}
	ErrNNotPositive   = AppError{Status: http.StatusBadRequest, Message: "n must be a positive integer"}
	ErrMNotPositive   = AppError{Status: http.StatusBadRequest, Message: "m must be a positive integer"}
	ErrMGreaterThanN  = AppError{Status: http.StatusBadRequest, Message: "m cannot be greater than n"}
	ErrKeyGeneration  = AppError{Status: http.StatusInternalServerError, Message: "Failed to generate key"}
	ErrKeyMarshalling = AppError{Status: http.StatusInternalServerError, Message: "Failed to marshal public key"}
)

func NewAppError(status int, message string) AppError {
	return AppError{
		Status:  status,
		Message: message,
	}
}

func WrapError(err error, status int, message string) AppError {
	return AppError{
		Status:  status,
		Message: fmt.Sprintf("%s: %v", message, err),
	}
}
