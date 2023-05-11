package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the URL Shortener API",
	})
}
