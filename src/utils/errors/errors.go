package errors

import (
	"net/http"
	"strings"
)

type RestError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message, detail string) *RestError {
	return &RestError{
		Message: message,
		Detail:  detail,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message, detail string) *RestError {
	return &RestError{
		Message: message,
		Detail:  detail,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message, detail string) *RestError {
	return &RestError{
		Message: message,
		Detail:  detail,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewStatusForbiddenError(message, detail string) *RestError {
	return &RestError{
		Message: message,
		Detail:  detail,
		Status:  http.StatusForbidden,
		Error:   "forbidden_request",
	}
}

func NewUnauthorizedError(message, detail string) *RestError {
	return &RestError{
		Message: message,
		Detail: detail,
		Status: http.StatusUnauthorized,
		Error: "unauthorized_request",
	}
}

func NewForbiddenError(message, detail string) *RestError {
	return &RestError{
		Message: message,
		Detail: detail,
		Status: http.StatusForbidden,
		Error: "forbidden",
	}
}

func CreateErrorMessageFromValidatorErrorList(err map[string]interface{}) string {
	var errRes []string
	for key, val := range err {
		errRes = append(errRes, key + " " + val.(error).Error())
	}
	return strings.Join(errRes, ", ")
}