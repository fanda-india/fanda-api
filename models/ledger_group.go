package models

// LedgerGroup db model
type LedgerGroup struct {
	ID          ID           `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty"`
	Code        string       `gorm:"size:16;uniqueIndex" json:"code,omitempty"`
	Name        string       `gorm:"size:25;uniqueIndex" json:"name,omitempty"`
	Description *string      `gorm:"size:255" json:"description,omitempty"`
	GroupType   byte         `gorm:"default:NULL" json:"groupType,omitempty"`
	ParentID    *ID          `gorm:"default:NULL" json:"parentId,omitempty"`
	Parent      *LedgerGroup `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"parent,omitempty"`
}
