package database

import (
	"github.com/softclub-go-0-0/url-shortener/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		&models.Link{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
