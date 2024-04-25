package main

import (
	"context"
	gerrors "errors"
	"fmt"
	"github.com/DarkMiMolle/Fiche/backend/controller"
	"github.com/DarkMiMolle/Fiche/backend/env"
	"github.com/DarkMiMolle/Fiche/backend/errors"
	"github.com/DarkMiMolle/Fiche/backend/middleware"
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

	mongoUri := os.Getenv(env.MongodbUri)
	fmt.Println(mongoUri)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)))
	if err != nil {
		panic(err.Error())
	}
	defer client.Disconnect(context.Background())

	databaseName := os.Getenv(env.DbName)
	db := client.Database(databaseName)
	server := gin.Default()
	server.Use(utils.SetupContextMiddleware(db), errors.Handle)

	server.POST("/api/signup", controller.SignUp)
	server.POST("/api/login", controller.Login)
	server.POST("/api/refresh")
	server.GET("/api/collection", func(c *gin.Context) {
		groupName, exists := c.GetQuery("collection")
		if !exists {
			panic(errors.HttpBadRequest{fmt.Errorf("missing 'collection' query parameter")})
		}

		result := db.Collection(os.Getenv(env.GroupCollection)).FindOne(c, bson.M{"user": "florent.carrez@yahoo.fr", "name": groupName})
		if result.Err() != nil && gerrors.Is(result.Err(), mongo.ErrNoDocuments) {
			c.JSON(utils.Success([]struct{}{}))
			return
		}
		if result.Err() != nil {
			fmt.Fprintf(os.Stderr, "%v\n", result.Err().Error())
			c.JSON(utils.InternalError(result.Err()))
			return
		}

		var coll models.Group
		if err := result.Decode(&coll); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", result.Err().Error())
			c.JSON(utils.InternalError(err))
			return
		}
		c.JSON(http.StatusOK, coll)
	})

	requiredAuth := server.Group("", middleware.JwtAuthMiddleware())

	requiredAuth.GET("/api/collections", controller.ListGroups)
	port := "3030"
	if val, ok := os.LookupEnv(env.ExposedPort); ok {
		port = val
	}
	server.Run("0.0.0.0:" + port)
}
