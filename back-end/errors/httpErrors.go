package errors

import "net/http"

type ErrorValue = error

type httpError interface {
	error
	Status() int
	httpError()
}

type HttpInternalServer struct {
	ErrorValue
}

func (HttpInternalServer) Status() int {
	return http.StatusInternalServerError
}

func (HttpInternalServer) httpError() {}

type HttpBadRequest struct {
	ErrorValue
}

func (HttpBadRequest) Status() int {
	return http.StatusBadRequest
}

func (HttpBadRequest) httpError() {}

type HttpUnauthenticatedError struct {
	ErrorValue
}

func (HttpUnauthenticatedError) Status() int {
	return http.StatusUnauthorized
}

func (HttpUnauthenticatedError) httpError() {}
