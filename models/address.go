package models

// Address db model
type Address struct {
	ID         uint    `gorm:"primaryKey;autoIncrement;not null"`
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
