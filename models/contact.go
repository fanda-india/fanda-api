package models

// Contact model
type Contact struct {
	Base
	Salutation   *string
	FirstName    *string
	LastName     *string
	Designation  *string
	Department   *string
	Email        *string
	WorkPhone    *string
	MobileNumber *string
}
