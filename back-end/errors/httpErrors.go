package errors

import (
	"fmt"
	"net/http"
)

func appCodeFrom(err error, or ...int) int {
	if appErr, isAppError := err.(App); isAppError {
		return appErr.Code
	}
	if len(or) > 0 {
		return or[0]
	}
	return 0
}

type Http struct {
	AppCode int    `json:"appCode"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

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
	unauthorised = (0b01 << iota) * 100
	tokenExpired
)
const (
	invalidPseudo = (0b01 << iota) * 1000
)

func InvalidPseudo() Http {
	return Http{
		AppCode: invalidPseudo,
		Status:  http.StatusBadRequest,
		Message: "le pseudo n'est pas valide",
	}
}

func TokenExpired() Http {
	return Http{
		AppCode: tokenExpired,
		Status:  http.StatusUnauthorized,
		Message: "le token d'authentication a expiré",
	}
}

func Unauthorised(reason error) Http {
	return Http{
		AppCode: appCodeFrom(reason, unauthorised),
		Status:  http.StatusUnauthorized,
		Message: fmt.Sprintf("unauthorised: %v", reason.Error()),
	}
}

func Internal(err error) Http {
	return Http{
		Status:  http.StatusInternalServerError,
		AppCode: appCodeFrom(err),
		Message: err.Error(),
	}
}
