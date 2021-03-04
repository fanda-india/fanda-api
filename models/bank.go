package models

// Bank db model
type Bank struct {
	ID       ID `gorm:"primaryKey;autoIncrement;not null"`
	LedgerID ID
	Ledger   Ledger `gorm:"foreignKey:LedgerID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
}
