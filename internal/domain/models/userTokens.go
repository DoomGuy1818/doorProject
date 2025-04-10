package models

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model
	ExpiredAt time.Time `json:"expired_at"`
	WorkerID  uint      `json:"user_id"`
	TokenHash string    `json:"token_hash"`
	IsValid   bool      `json:"is_valid" default:"true"`
}
