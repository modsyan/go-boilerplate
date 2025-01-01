package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	var httpError HttpError
	if errors.As(err, &httpError) {
		c.JSON(httpError.StatusCode(), gin.H{
			"message": httpError.Message(),
			"error":   httpError.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
