package middleware

import (
	"github.com/DarkMiMolle/Fiche/backend/errors"
	"github.com/gin-gonic/gin"
)

func HandleError() func(c *gin.Context) {
	return func(c *gin.Context) {
		defer errors.Recover(c)
		c.Next()
	}
}
