package models

import "fanda-api/enums"

// LedgerGroup db model
type LedgerGroup struct {
	ID          enums.ID     `gorm:"primaryKey;autoIncrement;not null"`
	Code        string       `gorm:"size:16;uniqueIndex"`
	Name        string       `gorm:"size:25;uniqueIndex"`
	Description *string      `gorm:"size:255"`
	GroupType   byte         `gorm:"default:NULL"`
	ParentID    *enums.ID    `gorm:"default:NULL"`
	Parent      *LedgerGroup `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
}
