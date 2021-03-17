package dtos

import (
	"fanda-api/enums"
	"fanda-api/models"
)

// BankDto model
type BankDto struct {
	ID models.ID `json:"id"`
	// LedgerID      models.ID              `json:"ledgerId"`
	// Code          string                 `json:"code,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Description   *string                `json:"description,omitempty"`
	AccountNumber *string                `json:"accountNumber,omitempty"`
	AccountType   *enums.BankAccountType `json:"accountType,omitempty"`
	IfscCode      *string                `json:"ifscCode,omitempty"`
	MicrCode      *string                `json:"micrCode,omitempty"`
	BranchCode    *string                `json:"branchCode,omitempty"`
	BranchName    *string                `json:"branchName,omitempty"`
	Address       *models.Address        `json:"address,omitempty"`
	Contact       *models.Contact        `json:"contact,omitempty"`
	IsDefault     *bool                  `json:"isDefault,omitempty"`
	IsActive      *bool                  `json:"isActive,omitempty"`
}

func (b *BankDto) ToBank() *models.Bank {
	return &models.Bank{
		ID:       b.ID,
		LedgerID: 0,
		Ledger: &models.Ledger{
			Name: b.Name, Description: b.Description,
			LedgerType: enums.Bank, IsActive: b.IsActive,
		},
		AccountNumber: b.AccountNumber,
		AccountType:   b.AccountType,
		IfscCode:      b.IfscCode,
		MicrCode:      b.MicrCode,
		BranchCode:    b.BranchCode,
		BranchName:    b.BranchName,
		Address:       b.Address,
		Contact:       b.Contact,
		IsDefault:     b.IsDefault,
	}
}

func (b *BankDto) FromBank(bank models.Bank) *BankDto {
	b.ID = bank.ID
	// b.Code = bank.Ledger.Code
	b.Name = bank.Ledger.Name
	b.Description = bank.Ledger.Description
	b.IsActive = bank.Ledger.IsActive
	b.AccountNumber = bank.AccountNumber
	b.AccountType = bank.AccountType
	b.IfscCode = bank.IfscCode
	b.MicrCode = bank.MicrCode
	b.BranchCode = bank.BranchCode
	b.BranchName = bank.BranchName
	b.Address = bank.Address
	b.Contact = bank.Contact
	b.IsDefault = bank.IsDefault

	return b
}
