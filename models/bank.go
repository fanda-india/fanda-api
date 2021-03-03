package models

import "fanda-api/enums"

// Bank db model
type Bank struct {
	ID       enums.ID `gorm:"primaryKey;autoIncrement;not null"`
	LedgerID enums.ID
	Ledger   Ledger `gorm:"foreignKey:LedgerID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
}
