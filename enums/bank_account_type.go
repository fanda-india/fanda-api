package enums

// BankAccountType enum
type BankAccountType byte

const (
	// Default bank account
	Default = iota
	// Savings bank account
	Savings
	// Current bank account
	Current
	// Fixed bank account
	Fixed
	// Demat bank account
	Demat
	// Salary bank account
	Salary
	// Other bank account
	Other
)
