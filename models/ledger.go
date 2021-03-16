package models

import (
	"time"
)

// Ledger db model
type Ledger struct {
	ID           ID            `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty"`
	Code         string        `gorm:"size:16;index:idx_ledgers_code,unique" json:"code,omitempty"`
	Name         string        `gorm:"size:25;index:idx_ledgers_name,unique" json:"name,omitempty"`
	Description  *string       `gorm:"size:255" json:"description,omitempty"`
	GroupID      *ID           `gorm:"default:NULL" json:"groupId,omitempty"`
	Group        *LedgerGroup  `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"group,omitempty"`
	LedgerType   byte          `gorm:"default:NULL" json:"ledgerType,omitempty"`
	IsSystem     *bool         `gorm:"default:false" json:"isSystem,omitempty"`
	OrgID        OrgID         `gorm:"index:idx_ledgers_code,unique;index:idx_ledgers_name,unique" json:"orgId,omitempty"`
	Organization *Organization `gorm:"foreignKey:OrgID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"-"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
	UpdatedAt    time.Time     `json:"updatedAt,omitempty"`
	Active       *bool         `gorm:"default:true" json:"active,omitempty"`
}
