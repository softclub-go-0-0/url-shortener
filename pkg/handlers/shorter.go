package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/url-shortener/pkg/helpers"
	"github.com/softclub-go-0-0/url-shortener/pkg/models"
	"log"
	"net/http"
)

func (h *handler) CreateLink(c *gin.Context) {
	var link models.Link
	err := c.ShouldBindJSON(&link)
	if err != nil {
		helpers.StatusBadRequest(c, err)
		return
	}
	if err := h.DB.Where("long_url =?", link.LongURL).First(&link).Error; err != nil {
		link.ShortURL = helpers.RandStr(8)
		if h.DB.Create(&link).Error != nil {
			log.Println("Inserting link data to DB:", err)
			helpers.InternalServerError(c)
			return
		}
	}
	host := "http://localhost:9999/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "Short Url created successfully",
		"short_url": host + link.ShortURL,
	})
}

func (h *handler) Redirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	var link models.Link
	if err := h.DB.Where("short_url =?", shortUrl).First(&link).Error; err != nil {
		helpers.NotFound(c, link.ShortURL)
		return
	}
	c.Redirect(http.StatusMovedPermanently, link.LongURL)
}

func (h *handler) DeleteLink(c *gin.Context) {
	var redirect models.Link
	if err := h.DB.Where("short_url = ?", c.Param("shortUrl")).First(&redirect).Error; err != nil {
		log.Println("client error - cannot find redirect:", err)
		helpers.NotFound(c, redirect.ShortURL)
		return
	}

	if err := h.DB.Delete(&redirect).Error; err != nil {
		log.Println("failed to remove server issues", err)
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "redirect removed successfully",
	})
}
