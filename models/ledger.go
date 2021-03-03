package models

import (
	"fanda-api/enums"
	"time"
)

// Ledger db model
type Ledger struct {
	ID             enums.ID    `gorm:"primaryKey;autoIncrement;not null"`
	Code           string      `gorm:"size:16;index:idx_ledgers_code,unique"`
	Name           string      `gorm:"size:25;index:idx_ledgers_name,unique"`
	Description    *string     `gorm:"size:255"`
	GroupID        enums.ID    `gorm:"default:NULL"`
	Group          LedgerGroup `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
	LedgerType     byte        `gorm:"default:NULL"`
	IsSystem       bool        `gorm:"default:false"`
	OrganizationID enums.ID    `gorm:"index:idx_ledgers_code,unique;index:idx_ledgers_name,unique"`
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	Active         bool `gorm:"default:true"`
}
