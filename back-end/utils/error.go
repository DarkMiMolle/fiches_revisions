package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonErr(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}

func Success(ret any) (int, any) {
	return http.StatusOK, ret
}
