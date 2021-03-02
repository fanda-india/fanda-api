package models

// Organization model
type Organization struct {
	Base
	Code         string
	Name         string
	Description  *string
	RegdNum      *string
	PAN          *string
	TAN          *string
	GSTIN        *string
	AddressID    uint
	Address      Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ContactID    uint
	Contact      Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ActiveYearID *uint
	Audit        Audit `gorm:"embedded"`
	Active       bool  `gorm:"default:1"`
}
