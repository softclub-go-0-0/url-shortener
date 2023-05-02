package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UrlShorter struct {
	ID        uuid.UUID      `gorm:"primary_key" json:"id"`
	LongUrl   string         `json:"long_url" binding:"required,url"`
	ShortUrl  string         `json:"short_url"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (base *UrlShorter) BeforeCraete(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
