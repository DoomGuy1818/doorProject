package models

import (
	"time"
)

type Exception struct {
	Date        time.Time `gorm:"primaryKey;type:date" json:"date"`
	Type        string    `gorm:"size:50;not null" json:"type"`
	Description string    `gorm:"type:text" json:"description"`
}
