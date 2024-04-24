package utils

import (
	"net/http"
)

func Success(ret any) (int, any) {
	return http.StatusOK, ret
}
