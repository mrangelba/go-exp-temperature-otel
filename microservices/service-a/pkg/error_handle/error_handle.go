package error_handle

import (
	"errors"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type RestError interface {
	APIError() (int, any)
}

type errorRestError struct {
	status int
	err    error
}

func (e errorRestError) Error() string {
	return e.err.Error()
}

func (e errorRestError) APIError() (int, any) {
	return e.status, Response{
		Message: e.Error(),
	}
}

func Handle(err error) (int, any) {
	var apiErr RestError
	if errors.As(err, &apiErr) {
		return apiErr.APIError()
	}

	return http.StatusInternalServerError, Response{
		Message: err.Error(),
	}
}

var (
	ErrNotFound            = &errorRestError{status: http.StatusNotFound, err: errors.New("can not found zipcode")}
	ErrUnprocessableEntity = &errorRestError{status: http.StatusUnprocessableEntity, err: errors.New("invalid zipcode")}
)
