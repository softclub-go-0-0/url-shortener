package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"github.com/softclub-go-0-0/url-shortener/pkg/models"
	"github.com/softclub-go-0-0/url-shortener/pkg/shortener"
	"net/http"
)

func (h *handler) CreateQrcode(c *gin.Context) {
	var link models.UrlShorter
	err := c.ShouldBindJSON(&link)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	link.ShortUrl = shortener.RandStr(8)
	host := "localhost:9999/"
	content := c.DefaultQuery("content", host+link.ShortUrl)
	if pic, err := qrcode.Encode(content, qrcode.Medium, 256); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.Data(http.StatusOK, "image/png", pic)
	}

}
