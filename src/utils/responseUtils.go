package utils

import (
	"net/http"

	"segmenta/src/models"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendErrorResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, models.Response{
		Success: false,
		Message: message,
	})
}
