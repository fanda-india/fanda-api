package models

// Address model
type Address struct {
	Base
	Attention  *string
	Line1      *string
	Line2      *string
	City       *string
	State      *string
	Country    *string
	PostalCode *string
	Phone      *string
	Fax        *string
}
