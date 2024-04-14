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

func InternalError(err error) (int, gin.H) {
	return http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	}
}

func BadRequestError(err error) (int, gin.H) {
	return http.StatusBadRequest, gin.H{
		"message": err.Error(),
	}
}

func Success(ret any) (int, any) {
	return http.StatusOK, ret
}
