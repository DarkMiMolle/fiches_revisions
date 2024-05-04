package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func Handle(c *gin.Context) {
	defer func() {
		rec := recover()
		if rec == nil {
			return
		}
		switch rec := rec.(type) {
		case Http:
			c.JSON(rec.Status, rec)
		case App:
			fmt.Fprintf(os.Stderr, "ERROR (%v): %v at: %v", rec.Code, rec.Message, strings.ReplaceAll(rec.GetStackString(), "\n", "\n\t"))
			c.JSON(http.StatusInternalServerError, Http{
				AppCode: rec.Code,
				Status:  http.StatusInternalServerError,
				Message: rec.Message,
			})
		default:
			panic(rec)
		}
		c.Abort()
	}()
	c.Next()
}
