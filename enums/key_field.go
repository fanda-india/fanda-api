package enums

import "fanda-api/utils"

// KeyField enum
type KeyField uint8

const (
	// Code key field
	Code KeyField = iota + 1
	// Name key field
	Name
	// Email key field
	Email
	// Mobile key field
	Mobile
	// Number key field
	Number
)

var keyFields = []string{"", "Code", "Name", "Email", "Mobile", "Number"}

// String method
func (k KeyField) String() string {
	return keyFields[k]
}

// KeyFieldConst method
func KeyFieldConst(s string) KeyField {
	return KeyField(utils.Search(keyFields, s))
}
