package errors

import (
	"fmt"
	"net/http"
)

type ErrorValue = error

type Http struct {
	AppCode int    `json:"appCode"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

/*
type Http = struct {
	@[json.Fields: json.LowerSnakeCaseFields]
	AppCode, Status int
	Message string
} + extension {
	const operator (Self & another Self) Self {
		self.AppCode |= another.AppCode
		self.Message += "; \{another.Message}"
		return self
	}
}
*/

func (h Http) And(another Http) Http {
	h.AppCode = h.AppCode | another.AppCode
	h.Message += "; " + another.Message
	return h
}

func BadRequest(err error) Http {
	return Http{
		AppCode: 0,
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func NewBadRequest(format string, v ...any) Http {
	return BadRequest(fmt.Errorf(format, v...))
}

const (
	invalidPseudo = (0b01 << iota) * 1000
	unauthenticated
)

func InvalidPseudo() Http {
	return Http{
		AppCode: invalidPseudo,
		Status:  http.StatusBadRequest,
		Message: "le pseudo n'est pas valide",
	}
}

func Unauthenticated(msg string, v ...any) Http {
	if msg == "" {
		msg = "le token d'authantification est n'est pas valide (ou manquant)"
	}
	return UnauthenticatedWith(fmt.Errorf(msg, v...))
}

func UnauthenticatedWith(err error) Http {
	return Http{
		AppCode: unauthenticated,
		Status:  http.StatusUnauthorized,
		Message: err.Error(),
	}
}

func Internal(err error) Http {
	return Http{
		Status:  http.StatusInternalServerError,
		AppCode: 0,
		Message: err.Error(),
	}
}

type httpError interface {
	error
	Status() int
	httpError()
}
