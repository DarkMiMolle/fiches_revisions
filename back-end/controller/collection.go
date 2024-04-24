package controller

import (
	"github.com/DarkMiMolle/Fiche/backend/env"
	"github.com/DarkMiMolle/Fiche/backend/errors"
	"github.com/DarkMiMolle/Fiche/backend/models"
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func ListGroups(c *gin.Context) {
	token, err := utils.GetTokenInfo(c)
	if err != nil {
		panic(errors.Internal(err))
	}
	db := utils.ExtractValues(c)
	collection := db.Collection(os.Getenv(env.GroupCollection))
	result, err := collection.Find(c, models.User{Pseudo: token.UserPseudo}.MongoIDFilter())
	if err != nil {
		panic(errors.BadRequest(err))
	}
	var groups []models.Group
	if err := result.All(c, &groups); err != nil {
		panic(errors.Internal(err))
	}

	c.JSON(utils.Success(groups))
}
