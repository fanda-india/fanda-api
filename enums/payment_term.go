package enums

// PaymentTerm enum
type PaymentTerm byte

const (
	// Immediate enum
	Immediate = iota
	// OnReceipt enum
	OnReceipt
	// OnDate enum
	OnDate
	// Net7 enum
	Net7
	// Net10 enum
	Net10
	// Net30 enum
	Net30
	// Net45 enum
	Net45
	// Net60 enum
	Net60
	// Net90 enum
	Net90
)
