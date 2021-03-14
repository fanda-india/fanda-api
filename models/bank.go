package models

import (
	"fanda-api/enums"
)

// Bank db model
type Bank struct {
	ID            ID                     `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty"`
	LedgerID      ID                     `json:"ledgerId,omitempty"`
	Ledger        *Ledger                `gorm:"foreignKey:LedgerID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"ledger"`
	AccountNumber *string                `gorm:"size:25;uniqueIndex" json:"accountNumber,omitempty"`
	AccountType   *enums.BankAccountType `json:"accountType,omitempty"`
	IfscCode      *string                `gorm:"size:16" json:"ifscCode,omitempty"`
	MicrCode      *string                `gorm:"size:16" json:"micrCode,omitempty"`
	BranchCode    *string                `gorm:"size:16" json:"branchCode,omitempty"`
	BranchName    *string                `gorm:"size:26" json:"branchName"`
	AddressID     ID                     `json:"addressId"`
	Address       *Address               `gorm:"foreignKey:AddressID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"address"`
	ContactID     ID                     `json:"contactId"`
	Contact       Contact                `gorm:"foreignKey:ContactID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"contact"`
	IsDefault     *bool                  `json:"isDefault"`
}
