package errors

import (
	"fmt"
	"github.com/DarkMiMolle/Fiche/backend/models"
	"github.com/gin-gonic/gin"
	"strings"
)

func Handle(c *gin.Context) {
	defer func() {
		rec := recover()
		if rec == nil {
			return
		}
		switch rec := rec.(type) {
		case httpError:
			c.JSON(rec.Status(), gin.H{"code": rec.Status(), "message": rec.Error()})
		case models.Error:
			fmt.Printf("ERROR (%v): %v\n\tat: %v\n", rec.Code, rec.Message, strings.Join(rec.GetStackString(), "|\n"))
			c.JSON(int(rec.Code), rec)
		default:
			panic(rec)
		}
		c.Abort()
	}()
	c.Next()
}
