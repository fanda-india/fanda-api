package models

// LedgerGroup db model
type LedgerGroup struct {
	ID          ID           `gorm:"primaryKey;autoIncrement;not null"`
	Code        string       `gorm:"size:16;uniqueIndex"`
	Name        string       `gorm:"size:25;uniqueIndex"`
	Description *string      `gorm:"size:255"`
	GroupType   byte         `gorm:"default:NULL"`
	ParentID    *ID          `gorm:"default:NULL"`
	Parent      *LedgerGroup `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
}
