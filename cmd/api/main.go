package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/url-shortener/pkg/handlers"
	"github.com/softclub-go-0-0/url-shortener/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func DBInit(user, password, dbname, port string) (*gorm.DB, error) {
	dsn := "host=localhost" +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable" +
		" TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&models.UrlShorter{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	DBName := flag.String("dbname", "redirect_db", "Enter the name of DB")
	DBUser := flag.String("dbuser", "postgres", "Enter the name of a DB user")
	DBPassword := flag.String("dbpassword", "85245655vG", "Enter the password of user")
	DBPort := flag.String("dbport", "5432", "Enter the port of DB")
	flag.Parse()

	db, err := DBInit(*DBUser, *DBPassword, *DBName, *DBPort)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	log.Println("Successfully connected to DB")

	h := handlers.NewHandler(db)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	router.POST("/create-short-url", h.CreateShortUrl)
	router.GET("/:shortUrl", h.HandlerShortUrlRedirect)
	router.DELETE("/:shortUrl", h.DeleteRedirectURL)

	err = router.Run(":9999")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
