package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/url-shortener/pkg/auth"
	"github.com/softclub-go-0-0/url-shortener/pkg/helpers"
	"github.com/softclub-go-0-0/url-shortener/pkg/models"
	"log"
	"net/http"
	"os"
)

type LinkData struct {
	LongURL string `json:"long_url" binding:"required,url"`
}

func (h *handler) CreateLink(c *gin.Context) {
	userData, exist := c.Get("user")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user := userData.(*auth.User)

	var linkData LinkData
	err := c.ShouldBindJSON(&linkData)
	if err != nil {
		helpers.StatusBadRequest(c, err)
		return
	}

	var link models.Link
	link.UserID = user.Id
	link.LongURL = linkData.LongURL
	link.ShortURL = helpers.RandStr(8)

	if h.DB.Create(&link).Error != nil {
		log.Println("Inserting link data to DB:", err)
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Short Url created successfully",
		"short_url": os.Getenv("APP_URL") + link.ShortURL,
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
	userData, exist := c.Get("user")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user := userData.(*auth.User)

	var link models.Link
	if err := h.DB.Where("short_url = ?", c.Param("shortUrl")).Where("user_id = ?", user.Id).First(&link).Error; err != nil {
		log.Println("client error - cannot find link:", err)
		helpers.NotFound(c, link.ShortURL)
		return
	}

	if err := h.DB.Delete(&link).Error; err != nil {
		log.Println("failed to remove server issues", err)
		helpers.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "link removed successfully",
	})
}
