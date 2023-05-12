package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/softclub-go-0-0/url-shortener/pkg/database"
	"github.com/softclub-go-0-0/url-shortener/pkg/handlers"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.DBInit(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	log.Println("Successfully connected to DB")

	h := handlers.NewHandler(db)

	router := gin.Default()
	router.GET("/", h.Welcome)

	router.POST("/links", h.CreateLink)
	router.GET("/:shortUrl", h.Redirect)
	router.POST("/qrcode", h.CreateQRCode)
	router.DELETE("/:shortUrl", h.DeleteLink)

	err = router.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
