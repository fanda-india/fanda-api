package enums

// KeyField enum
type KeyField uint8

const (
	// Code key field
	Code KeyField = iota
	// Name key field
	Name
	// Email key field
	Email
	// Mobile key field
	Mobile
	// Number key field
	Number
)

func (k KeyField) String() string {
	return [...]string{"Code", "Name", "Email", "Mobile", "Number"}[k]
}
