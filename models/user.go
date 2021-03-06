package models

import (
	"time"
)

// User db model
type User struct {
	ID              ID      `gorm:"primaryKey;autoIncrement;not null"`
	UserName        string  `gorm:"size:16;uniqueIndex"`
	Email           string  `gorm:"size:100;uniqueIndex"`
	MobileNumber    string  `gorm:"size:10;uniqueIndex"`
	FirstName       *string `gorm:"size:25"`
	LastName        *string `gorm:"size:25"`
	Password        string  `gorm:"size:100"`
	IsPasswordReset bool    `gorm:"default:false"`
	LoginAt         *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	IsActive        *bool `gorm:"default:true"`
}
