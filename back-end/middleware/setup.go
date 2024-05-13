package middleware

import (
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupContext(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(utils.ContextDB, db)
	}
}
