package main

import (
	"context"
	"fmt"
	"github.com/DarkMiMolle/Fiche/backend/models"
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello there !")

	mongoUri := os.Getenv("MONGODB_URI")
	fmt.Println(mongoUri)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err.Error())
	}
	defer client.Disconnect(context.Background())

	databaseName := os.Getenv("DB_NAME")
	db := client.Database(databaseName)
	server := gin.Default()
	server.Use(utils.SetupContextMiddleware(db))

	server.GET("/api/collection", func(c *gin.Context) {
		collection, exists := c.GetQuery("collection")
		if !exists {
			c.JSON(http.StatusBadRequest, "missing query 'collection'")
			return
		}

		result, err := db.Collection("collections").Find(c, bson.M{"user": "florent.carrez@yahoo.fr", "name": collection})
		if err != nil {
			c.JSON(http.StatusInternalServerError, c.Error(err))
			return
		}
		var coll []models.Collection
		if err := result.All(c, &coll); err != nil {
			c.JSON(http.StatusInternalServerError, c.Error(err))
			return
		}
		c.JSON(http.StatusOK, coll)
	})
	fmt.Println(os.Getenv("URL"))
	server.Run("0.0.0.0:3030")
}
