package utils

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

const ContextDB = "DB"

type SharedValue struct {
	Db *mongo.Database
}

func ExtractSharedValue(c *gin.Context) (val SharedValue) {
	val.Db = ExtractValues(c)
	return
}

func ExtractValues(c *gin.Context) (db *mongo.Database) {
	db = c.MustGet(ContextDB).(*mongo.Database)
	return
}
