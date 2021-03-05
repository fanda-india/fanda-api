package models

// Address db model
type Address struct {
	ID         ID      `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty"`
	Attention  *string `gorm:"size:25" json:"attention,omitempty"`
	Line1      *string `gorm:"size:50" json:"line1,omitempty"`
	Line2      *string `gorm:"size:50" json:"line2,omitempty"`
	City       *string `gorm:"size:25" json:"city,omitempty"`
	State      *string `gorm:"size:25" json:"state,omitempty"`
	Country    *string `gorm:"size:25" json:"country,omitempty"`
	PostalCode *string `gorm:"size:10" json:"postalCode,omitempty"`
	Phone      *string `gorm:"size:25" json:"phone,omitempty"`
	Fax        *string `gorm:"size:25" json:"fax,omitempty"`
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
