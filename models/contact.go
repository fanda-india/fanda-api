package models

// Contact db model
type Contact struct {
	ID           uint    `gorm:"primaryKey;autoIncrement;not null"`
	Salutation   *string `gorm:"size:5"`
	FirstName    *string `gorm:"size:25"`
	LastName     *string `gorm:"size:25"`
	Designation  *string `gorm:"size:25"`
	Department   *string `gorm:"size:25"`
	Email        *string `gorm:"size:100"`
	WorkPhone    *string `gorm:"size:25"`
	MobileNumber *string `gorm:"size:25"`
}
