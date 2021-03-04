package models

import (
	"time"
)

// Organization db model
type Organization struct {
	ID           ID       `gorm:"primaryKey;autoIncrement;not null"`
	Code         string   `gorm:"size:16;uniqueIndex"`
	Name         string   `gorm:"size:50;uniqueIndex"`
	Description  *string  `gorm:"size:255"`
	RegdNum      *string  `gorm:"size:25"`
	PAN          *string  `gorm:"size:25"`
	TAN          *string  `gorm:"size:25"`
	GSTIN        *string  `gorm:"size:25"`
	AddressID    *ID      `gorm:"default:NULL"`
	Address      *Address `gorm:"foreignKey:AddressID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
	ContactID    *ID      `gorm:"default:NULL"`
	Contact      *Contact `gorm:"foreignKey:ContactID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
	ActiveYearID *ID      `gorm:"default:NULL"`
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	Active       bool `gorm:"default:true"`
}
