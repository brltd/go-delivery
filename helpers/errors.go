package helpers

import "net/http"

type ApiError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e ApiError) AsMessage() *ApiError {
	return &ApiError{
		Message: e.Message,
	}
}

func NotFoundError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func InternalServerError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func ValidationError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func Unauthorized(message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}
