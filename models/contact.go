package models

// Contact db model
type Contact struct {
	ID           ID      `gorm:"primaryKey;autoIncrement;not null" json:"id,omitempty"`
	Salutation   *string `gorm:"size:5" json:"salutation,omitempty"`
	FirstName    *string `gorm:"size:25" json:"firstName,omitempty"`
	LastName     *string `gorm:"size:25" json:"lastName,omitempty"`
	Designation  *string `gorm:"size:25" json:"designation,omitempty"`
	Department   *string `gorm:"size:25" json:"department,omitempty"`
	Email        *string `gorm:"size:100" json:"email,omitempty"`
	WorkPhone    *string `gorm:"size:25" json:"workPhone,omitempty"`
	MobileNumber *string `gorm:"size:25" json:"mobileNumber,omitempty"`
}

// IsEmpty method
func (c *Contact) IsEmpty() bool {
	if c.Salutation == nil && c.FirstName == nil && c.LastName == nil &&
		c.Designation == nil && c.Department == nil && c.Email == nil &&
		c.WorkPhone == nil && c.MobileNumber == nil {
		return true
	}
	return false
}
