package enums

import "fanda-api/utils"

// KeyField enum
type KeyField byte

const (
	// IDField enum
	IDField KeyField = iota
	// CodeField enum
	CodeField
	// NameField enum
	NameField
	// EmailField enum
	EmailField
	// MobileField enum
	MobileField
	// NumberField enum
	NumberField
)

var keyFields = []string{"ID", "Code", "Name", "Email", "Mobile", "Number"}

// String method
func (k KeyField) String() string {
	return keyFields[k]
}

// KeyFieldConst method
func KeyFieldConst(s string) KeyField {
	return KeyField(utils.Search(keyFields, s))
}
