package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Link struct {
	ID        uuid.UUID      `gorm:"primary_key" json:"id"`
	LongUrl   string         `json:"long_url" binding:"required,url"`
	ShortUrl  string         `json:"short_url"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (b *Link) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
