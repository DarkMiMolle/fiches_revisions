package middleware

import (
	"fmt"
	"github.com/DarkMiMolle/Fiche/backend/errors"
	"github.com/DarkMiMolle/Fiche/backend/models"
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer errors.Handle(c)
		token, err := utils.GetToken(c)
		if err != nil {
			panic(errors.HttpUnauthenticatedError{err})
		}
		tokenInfo, err := utils.ExtractTokenInfo(token)
		if err != nil {
			panic(errors.HttpUnauthenticatedError{fmt.Errorf("unauthorized: %v", err.Error())})
		}
		if tokenInfo.ExpireDate.Before(time.Now()) {
			panic(errors.HttpUnauthenticatedError{models.AuthTokenExpiredError.Make()})
		}
		utils.AddToken(c, tokenInfo)
		c.Next()
	}
}
