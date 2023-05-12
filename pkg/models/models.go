package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Link struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	LongURL   string         `json:"long_url"`
	ShortURL  string         `json:"short_url"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
