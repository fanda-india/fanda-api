package enums

// LedgerGroupType enum
type LedgerGroupType uint8

const (
	// Asset group type
	Asset LedgerGroupType = iota + 1
	// Liability group type
	Liability
	// Revenue group type
	Revenue
	// Income group type
	Income
	// Expense group type
	Expense
	// Branch
	// Warehouse
)

func (l LedgerGroupType) String() string {
	return [...]string{"Asset", "Liability", "Revenue", "Income", "Expense"}[l]
}
