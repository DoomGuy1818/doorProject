package models

import (
	"doorProject/internal/domain/value"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	FullName string            `json:"full_name" validate:"required"`
	Phone    value.PhoneNumber `json:"phone" validate:"required"`
	Email    value.Email       `json:"email" validate:"required,email"`
	CartID   uint              `json:"cart_id" validate:"required"`
}
