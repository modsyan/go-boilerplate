package validators

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindJsonAndValidateRequest(c *gin.Context, req interface{}, validator IValidator) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON payload",
			"error":   err.Error(),
		})
		return false
	}

	validationErr := validator.ValidateStruct(req)
	if validationErr != nil {
		c.JSON(validationErr.StatusCode(), gin.H{
			"message": validationErr.Message(),
			"errors":  validationErr.ValidationErrors(),
		})
		return false
	}

	return true
}

func BindQueryAndValidateRequest(c *gin.Context, req interface{}, validator IValidator) bool {
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query parameters",
			"error":   err.Error(),
		})
		return false
	}

	validationErr := validator.ValidateStruct(req)
	if validationErr != nil {
		c.JSON(validationErr.StatusCode(), gin.H{
			"message": validationErr.Message(),
			"errors":  validationErr.ValidationErrors(),
		})
		return false
	}

	return true
}

func ValidateRequestOnly(c *gin.Context, req interface{}, validator IValidator) bool {
	validationErr := validator.ValidateStruct(req)
	if validationErr != nil {
		c.JSON(validationErr.StatusCode(), gin.H{
			"message": validationErr.Message(),
			"errors":  validationErr.ValidationErrors(),
		})
		return false
	}
	return true
}
