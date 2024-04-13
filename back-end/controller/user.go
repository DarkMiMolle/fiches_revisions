package controller

import (
	"github.com/DarkMiMolle/Fiche/backend/models"
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SingUp(c *gin.Context) {
	// get user from context
	var body struct {
		Email    models.Email
		Password string
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, utils.JsonErr(err))
		return
	}
	// check if user exists in db

	// hash password

	// register the user
}

var _ gin.HandlerFunc = SingUp
