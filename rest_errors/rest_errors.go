package rest_errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"code"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewRestError(message string, status int, error string, causes []interface{}) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Error:   error,
		Causes:  causes,
	}
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: "unable to retrieve user information from given access_token",
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
