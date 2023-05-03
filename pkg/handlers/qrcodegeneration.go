package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"github.com/softclub-go-0-0/url-shortener/pkg/models"
	"github.com/softclub-go-0-0/url-shortener/pkg/shortener"
	"net/http"
)

//func (h *handler) CreateQrcode(c *gin.Context) {
//	var link models.UrlShorter
//	err := c.ShouldBindJSON(&link)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	if err := h.DB.Where("long_url =?", link.LongUrl).First(&link).Error; err != nil {
//		link.ShortUrl = shortener.RandStr(8)
//		if h.DB.Create(&link).Error != nil {
//			log.Println("Inserting link data to DB:", err)
//			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
//				"message": "Interval Server Error",
//			})
//			return
//		}
//		linkForCreateqr := link.ShortUrl
//
//		qrCode, err := qr.Encode(linkForCreateqr, qr.L, qr.Auto)
//		qrCode, err = barcode.Scale(qrCode, 512, 512)
//		if err != nil {
//			log.Println("Error to create QRCODE", err)
//			return
//		}
//		png.Encode(c, qrCode)
//	}
//
//}

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
