package database

import (
	"gorm.io/gorm"
	"time"
)

type AuthToken struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	IsValid   bool      `json:"is_valid"`
	Token     string    `json:"token"`
}
