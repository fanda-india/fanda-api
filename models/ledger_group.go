package models

// LedgerGroup db model
type LedgerGroup struct {
	ID          uint         `gorm:"primaryKey;autoIncrement;not null"`
	Code        string       `gorm:"size:16;uniqueIndex"`
	Name        string       `gorm:"size:25;unique;uniqueIndex"`
	Description *string      `gorm:"size:255"`
	GroupType   int          `gorm:"default:NULL"`
	ParentID    *int         `gorm:"default:NULL"`
	Parent      *LedgerGroup `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
