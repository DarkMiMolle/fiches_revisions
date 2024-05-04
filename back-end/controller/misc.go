package controller

import (
	"github.com/DarkMiMolle/Fiche/backend/errors"
	"github.com/gin-gonic/gin"
)

func TODO(c *gin.Context) {
	panic(errors.App{
		Message: "TODO",
	}.Stack())
}
