package models

// Party db model
type Party struct {
	ID          ID       `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty"`
	LedgerID    ID       `json:"ledgerId,omitempty"`
	Ledger      Ledger   `gorm:"foreignKey:LedgerID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"ledger"`
	RegdNum     *string  `gorm:"size:25" json:"regdNum,omitempty"`
	PAN         *string  `gorm:"size:25" json:"pan,omitempty"`
	TAN         *string  `gorm:"size:25" json:"tan,omitempty"`
	GSTIN       *string  `gorm:"size:25" json:"gstin,omitempty"`
	PaymentTerm *byte    `json:"paymentTerm,omitempty"`
	CreditLimit *float32 `json:"creditLimit,omitempty"`
	AddressID   ID       `json:"addressId"`
	Address     *Address `gorm:"foreignKey:AddressID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"address"`
	ContactID   ID       `json:"contactId"`
	Contact     Contact  `gorm:"foreignKey:ContactID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"contact"`
	IsDefault   *bool    `json:"isDefault"`
}
