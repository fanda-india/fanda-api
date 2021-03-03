package models

// LedgerGroup db model
type LedgerGroup struct {
	ID          uint         `gorm:"primaryKey;autoIncrement;not null"`
	Code        string       `gorm:"size:16;unique;not null;uniqueIndex"`
	Name        string       `gorm:"size:25;unique;not null;uniqueIndex"`
	Description *string      `gorm:"size:255"`
	GroupType   int          `gorm:"not null"`
	ParentID    *int         `gorm:"default:null"`
	Parent      *LedgerGroup `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
