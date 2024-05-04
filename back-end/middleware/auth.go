package middleware

import (
	"github.com/DarkMiMolle/Fiche/backend/errors"
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer errors.Handle(c)
		token, err := utils.GetToken(c)
		if err != nil {
			panic(errors.UnauthenticatedWith(err))
		}
		tokenInfo, err := utils.ExtractTokenInfo(token)
		if err != nil {
			panic(errors.Unauthenticated("unauthorized: %v", err.Error()))
		}
		if tokenInfo.ExpireDate.Before(time.Now()) {
			panic(errors.Unauthenticated("le token d'authentication a expiré"))
		}
		utils.AddToken(c, tokenInfo)
		c.Next()
	}
}
