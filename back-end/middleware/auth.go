package middleware

import (
	"time"

	"github.com/DarkMiMolle/Fiche/backend/errors"
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer errors.Recover(c)
		token, err := utils.GetToken(c)
		if err != nil {
			panic(errors.Unauthorised(err))
		}
		tokenInfo, err := utils.ExtractTokenInfo(token)
		if err != nil {
			panic(errors.Unauthorised(err))
		}
		if tokenInfo.ExpireDate.Before(time.Now()) {
			panic(errors.TokenExpired())
		}
		utils.AddToken(c, tokenInfo)
		c.Next()
	}
}
