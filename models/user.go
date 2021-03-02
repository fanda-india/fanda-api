package models

import "time"

// User db model
type User struct {
	Base
	UserName        string
	Email           string
	MobileNumber    string
	FirstName       string
	LastName        string
	Password        string
	IsPasswordReset bool `gorm:"default:false"`
	LoginAt         *time.Time
	Audit           Audit `gorm:"embedded"`
	Active          bool  `gorm:"default:1"`
}
