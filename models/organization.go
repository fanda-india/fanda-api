package models

import (
	"time"
)

// Organization db model
type Organization struct {
	ID           ID        `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty"`
	Code         string    `gorm:"size:16;uniqueIndex" json:"code,omitempty"`
	Name         string    `gorm:"size:50;uniqueIndex" json:"name,omitempty"`
	Description  *string   `gorm:"size:255" json:"description,omitempty"`
	RegdNum      *string   `gorm:"size:25" json:"regdNum,omitempty"`
	PAN          *string   `gorm:"size:25" json:"pan,omitempty"`
	TAN          *string   `gorm:"size:25" json:"tan,omitempty"`
	GSTIN        *string   `gorm:"size:25" json:"gstin,omitempty"`
	AddressID    *ID       `gorm:"default:NULL" json:"addressId,omitempty"`
	Address      *Address  `gorm:"foreignKey:AddressID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"address,omitempty"`
	ContactID    *ID       `gorm:"default:NULL" json:"contactId,omitempty"`
	Contact      *Contact  `gorm:"foreignKey:ContactID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"contact,omitempty"`
	ActiveYearID *ID       `gorm:"default:NULL" json:"activeYearId,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	Active       bool      `gorm:"default:true" json:"active,omitempty"`
}
