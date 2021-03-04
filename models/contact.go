package models

// Contact db model
type Contact struct {
	ID           ID      `gorm:"primaryKey;autoIncrement;not null"`
	Salutation   *string `gorm:"size:5"`
	FirstName    *string `gorm:"size:25"`
	LastName     *string `gorm:"size:25"`
	Designation  *string `gorm:"size:25"`
	Department   *string `gorm:"size:25"`
	Email        *string `gorm:"size:100"`
	WorkPhone    *string `gorm:"size:25"`
	MobileNumber *string `gorm:"size:25"`
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
