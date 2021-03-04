package enums

import (
	"strings"
)

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
	return KeyField(find(keyFields, s))
}

func find(a []string, x string) int {
	for i, n := range a {
		if strings.EqualFold(x, n) {
			return i
		}
	}
	return len(a)
}
