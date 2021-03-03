package models

import "time"

// Organization db model
type Organization struct {
	ID           uint     `gorm:"primaryKey;autoIncrement;not null"`
	Code         string   `gorm:"size:16;unique;not null;uniqueIndex"`
	Name         string   `gorm:"size:50;unique;not null;uniqueIndex"`
	Description  *string  `gorm:"size:255"`
	RegdNum      *string  `gorm:"size:25"`
	PAN          *string  `gorm:"size:25"`
	TAN          *string  `gorm:"size:25"`
	GSTIN        *string  `gorm:"size:25"`
	AddressID    *uint    `gorm:"default:null"`
	Address      *Address `gorm:"foreignKey:AddressID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ContactID    *uint    `gorm:"default:null"`
	Contact      *Contact `gorm:"foreignKey:ContactID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ActiveYearID *uint    `gorm:"default:null"`
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	Active       bool `gorm:"not null;default:true"`
}
