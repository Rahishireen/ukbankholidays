package utils

import "net/http"

type AppErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

// NewBadRequestError to send error for invalid input request
func NewBadRequestError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError to send error for data which is not present in Database
func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalServerError to send error for any unknown Internal Errors
func NewInternalServerError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "Internal_server_error",
	}
}
