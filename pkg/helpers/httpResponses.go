package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IntervalServerError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "Interval server error",
	})
}

func NotFound(c *gin.Context, model string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": model + "not found",
	})
}
func UnprocessableEntity(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "validation error",
		"errors":  err.Error(),
	})
}
func StatusBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": "I can't find such links!",
		"error":   err.Error(),
	})
}
