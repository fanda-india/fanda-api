package models

import "time"

// User db model
type User struct {
	ID              uint    `gorm:"primaryKey;autoIncrement;not null"`
	UserName        string  `gorm:"size:16;unique;not null;uniqueIndex"`
	Email           string  `gorm:"size:100;unique;not null;uniqueIndex"`
	MobileNumber    string  `gorm:"size:10;unique;not null;uniqueIndex"`
	FirstName       *string `gorm:"size:25"`
	LastName        *string `gorm:"size:25"`
	Password        string  `gorm:"size:100;not null"`
	IsPasswordReset bool    `gorm:"not null;default:false"`
	LoginAt         *time.Time
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	Active          bool `gorm:"not null;default:true"`
}
