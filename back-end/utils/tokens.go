package utils

import (
	"fmt"
	"github.com/DarkMiMolle/Fiche/backend/env"
	"github.com/DarkMiMolle/Fiche/backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

type TokenInfo struct {
	UserEmail  models.Email
	ExpireDate time.Time
}

func GenerateJwt(user *models.User) (string, error) {
	tokenDuration := 6 * time.Hour

	claims := jwt.MapClaims{
		"authorized": true,
		"user_email": user.Email,
		"exp":        time.Now().Add(tokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv(env.ApiSecret)))
}

func GetToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == "" {
		return "", fmt.Errorf("missing required Authorisation header")
	}
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", fmt.Errorf("invalid bearer token")
}

func ExtractTokenInfo(tokenStr string) (TokenInfo, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv(env.ApiSecret)), nil
	})
	if err != nil {
		return TokenInfo{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	info := TokenInfo{}
	if ok && token.Valid {
		info.UserEmail = models.Email(claims["user_email"].(string))
		info.ExpireDate = time.Unix(claims["exp"].(int64), 0)
		return info, nil
	}
	return info, fmt.Errorf("invalid token")
}

const contextToken = "Deduced Token"

func AddToken(c *gin.Context, token TokenInfo) {
	c.Set(contextToken, token)
}

//func GetToken(c *gin.Context) TokenInfo {
//	token, exists := c.Get(contextToken)
//	if !exists {
//		panic("GetToken must be call only if the appropriate middleware is call")
//	}
//	return token.(TokenInfo)
//}
