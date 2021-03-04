package models

// Address db model
type Address struct {
	ID         ID      `gorm:"primaryKey;autoIncrement;not null"`
	Attention  *string `gorm:"size:25"`
	Line1      *string `gorm:"size:50"`
	Line2      *string `gorm:"size:50"`
	City       *string `gorm:"size:25"`
	State      *string `gorm:"size:25"`
	Country    *string `gorm:"size:25"`
	PostalCode *string `gorm:"size:10"`
	Phone      *string `gorm:"size:25"`
	Fax        *string `gorm:"size:25"`
}

// IsEmpty method
func (a *Address) IsEmpty() bool {
	if a.Attention == nil && a.Line1 == nil && a.Line2 == nil &&
		a.City == nil && a.State == nil && a.Country == nil &&
		a.PostalCode == nil && a.Phone == nil && a.Fax == nil {
		return true
	}
	return false
}
