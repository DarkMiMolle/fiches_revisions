package errors

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
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
}
