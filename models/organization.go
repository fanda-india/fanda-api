package models

import (
	"fanda-api/enums"
	"time"
)

// Organization db model
type Organization struct {
	ID           enums.ID  `gorm:"primaryKey;autoIncrement;not null"`
	Code         string    `gorm:"size:16;uniqueIndex"`
	Name         string    `gorm:"size:50;uniqueIndex"`
	Description  *string   `gorm:"size:255"`
	RegdNum      *string   `gorm:"size:25"`
	PAN          *string   `gorm:"size:25"`
	TAN          *string   `gorm:"size:25"`
	GSTIN        *string   `gorm:"size:25"`
	AddressID    *enums.ID `gorm:"default:NULL"`
	Address      *Address  `gorm:"foreignKey:AddressID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
	ContactID    *enums.ID `gorm:"default:NULL"`
	Contact      *Contact  `gorm:"foreignKey:ContactID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
	ActiveYearID *enums.ID `gorm:"default:NULL"`
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	Active       bool `gorm:"default:true"`
}
