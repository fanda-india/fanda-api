package dtos

import (
	"fanda-api/enums"
	"fanda-api/models"
)

// BankDto model
type BankDto struct {
	ID models.ID `json:"id"`
	// LedgerID      models.ID              `json:"ledgerId"`
	Code          string                 `json:"code"`
	Name          string                 `json:"name"`
	Description   *string                `json:"description"`
	AccountNumber *string                `json:"accountNumber"`
	AccountType   *enums.BankAccountType `json:"accountType"`
	IfscCode      *string                `json:"ifscCode"`
	MicrCode      *string                `json:"micrCode"`
	BranchCode    *string                `json:"branchCode"`
	BranchName    *string                `json:"branchName"`
	// Address       *models.Address        `json:"address"`
	// Contact       *models.Contact        `json:"contact"`
	IsDefault *bool `json:"isDefault"`
	Active    *bool `json:"active"`
}

func (b *BankDto) ToBank() *models.Bank {
	return &models.Bank{
		ID:            b.ID,
		LedgerID:      0,
		Ledger:        &models.Ledger{Code: b.Code, Name: b.Name, Description: b.Description, LedgerType: enums.Bank, Active: b.Active},
		AccountNumber: b.AccountNumber,
		AccountType:   b.AccountType,
		IfscCode:      b.IfscCode,
		MicrCode:      b.MicrCode,
		BranchCode:    b.BranchCode,
		BranchName:    b.BranchName,
		// AddressID:     b.Address.ID,
		// Address:       b.Address,
		// ContactID:     b.Contact.ID,
		// Contact:       *b.Contact,
		IsDefault: b.IsDefault,
	}
}

func (b *BankDto) FromBank(bank *models.Bank) *models.Bank {
	b.ID = bank.ID
	b.Code = bank.Ledger.Code
	b.Name = bank.Ledger.Name
	b.Description = bank.Ledger.Description
	b.Active = bank.Ledger.Active
	b.AccountNumber = bank.AccountNumber
	b.AccountType = bank.AccountType
	b.IfscCode = bank.IfscCode
	b.MicrCode = bank.MicrCode
	b.BranchCode = bank.BranchCode
	b.BranchName = bank.BranchName
	// b.Address = &models.Address{ID: bank.a}
	// b.Contact = &models.Contact{}
	b.IsDefault = bank.IsDefault

	return bank
}
