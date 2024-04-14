package controller

import (
	"github.com/DarkMiMolle/Fiche/backend/env"
	"github.com/DarkMiMolle/Fiche/backend/models"
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/DarkMiMolle/GTL/array"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func ListGroups(c *gin.Context) {
	token, err := utils.GetTokenInfo(c)
	if err != nil {
		c.JSON(utils.InternalError(err))
		return
	}
	db := utils.ExtractValues(c)
	collection := db.Collection(os.Getenv(env.GroupCollection))
	result, err := collection.Find(c, bson.M{
		"user": token.UserEmail,
	})
	if err != nil {
		c.JSON(utils.BadRequestError(err))
		return
	}
	var groups []models.Group
	if err := result.All(c, &groups); err != nil {
		c.JSON(utils.InternalError(err))
		return
	}

	c.JSON(utils.Success(array.Map(groups, func(elem models.Group) string {
		return elem.Name
	})))
}
