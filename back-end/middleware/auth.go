package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.GetToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.JsonErr(err))
			c.Abort()
			return
		}
		tokenInfo, err := utils.ExtractTokenInfo(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.JsonErr(fmt.Errorf("unauthorized: %v", err.Error())))
			c.Abort()
			return
		}
		if tokenInfo.ExpireDate.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, utils.JsonErr(fmt.Errorf("token expired")))
			c.Abort()
			return
		}
		utils.AddToken(c, tokenInfo)
		c.Next()
	}
}
