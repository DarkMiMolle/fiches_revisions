package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/DarkMiMolle/Fiche/backend/env"
	"github.com/DarkMiMolle/Fiche/backend/models"
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/DarkMiMolle/GTL/optional"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

func SingUp(c *gin.Context) {
	// get user from context
	var body struct {
		Email    models.Email           `json:"email" binding:""`
		Password string                 `json:"password"`
		Name     optional.Value[string] `json:"name"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, utils.JsonErr(err))
		return
	}
	if !body.Email.IsValid() {
		c.JSON(http.StatusBadRequest, utils.JsonErr(fmt.Errorf("email '%v' is not valid", body.Email)))
		return
	}
	if len(body.Password) < 5 {
		c.JSON(http.StatusBadRequest, utils.JsonErr(fmt.Errorf("invalid password, requires at least 5 character")))
		return
	}

	// check if user email exists in db
	db := utils.ExtractValues(c)
	collection := db.Collection(os.Getenv(env.UserCollection))
	result := collection.FindOne(context.Background(), bson.M{
		"email": body.Email,
	})
	if result.Err() == nil {
		c.JSON(http.StatusBadRequest, utils.JsonErr(fmt.Errorf("user with same email (%v) already exists", body.Email)))
		return
	}
	if !errors.Is(result.Err(), mongo.ErrNoDocuments) {
		c.JSON(http.StatusInternalServerError, utils.JsonErr(result.Err()))
		return
	}

	// hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		c.JSON(http.StatusBadRequest, utils.JsonErr(err))
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.JsonErr(err))
	}

	// register the user
	user := models.User{
		Email:    body.Email,
		Name:     body.Name.ValueOr(""),
		Password: models.Password(hashPassword),
	}
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.JsonErr(err))
		return
	}

	c.JSON(utils.Success(user))
}
