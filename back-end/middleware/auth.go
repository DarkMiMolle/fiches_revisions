package middleware

import (
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		utils.AddToken(c, tokenInfo)
		c.Next()
	}
}
