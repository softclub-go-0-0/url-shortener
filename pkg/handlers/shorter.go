package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/url-shortener/pkg/models"
	"github.com/softclub-go-0-0/url-shortener/pkg/shortener"
	"log"
	"net/http"
)

func (h *handler) CreateShortUrl(c *gin.Context) {
	var link models.UrlShorter
	err := c.ShouldBindJSON(&link)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Where("long_url =?", link.LongUrl).First(&link).Error; err != nil {
		link.ShortUrl = shortener.RandStr(8)
		if h.DB.Create(&link).Error != nil {
			log.Println("Inserting link data to DB:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
	}
	host := "http://localhost:9999/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "Short Url created successfully",
		"short_url": host + link.ShortUrl,
	})
}

func (h *handler) HandlerShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	var link models.UrlShorter
	if err := h.DB.Where("short_url =?", shortUrl).First(&link).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "I can't find such links!",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, link.LongUrl)
}

func (h *handler) DeleteRedirectURL(c *gin.Context) {
	var redirect models.UrlShorter
	if err := h.DB.Where("short_url = ?", c.Param("shortUrl")).First(&redirect).Error; err != nil {
		log.Println("client error - cannot find redirect:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "I can't find such links!",
		})
		return
	}

	if err := h.DB.Delete(&redirect).Error; err != nil {
		log.Println("failed to remove server issues", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "redirect removed successfully",
	})
}
