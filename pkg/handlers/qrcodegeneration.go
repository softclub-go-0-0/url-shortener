package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"github.com/softclub-go-0-0/url-shortener/pkg/helpers"
	"github.com/softclub-go-0-0/url-shortener/pkg/models"
	"log"
	"net/http"
)

func (h *handler) CreateQRCode(c *gin.Context) {
	var link models.Link
	err := c.ShouldBindJSON(&link)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Where("long_url =?", link.LongURL).First(&link).Error; err != nil {
		link.ShortURL = helpers.RandStr(8)
		if h.DB.Create(&link).Error != nil {
			log.Println("Inserting link data to DB:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
	}
	host := "localhost:9999/"
	content := c.DefaultQuery("content", host+link.ShortURL)
	if pic, err := qrcode.Encode(content, qrcode.Medium, 256); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.Data(http.StatusOK, "image/png", pic)
	}
}
