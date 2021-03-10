package enums

// LedgerGroupType enum
type LedgerGroupType byte

const (
	// AssetGroup enum
	AssetGroup LedgerGroupType = iota + 1
	// LiabilityGroup enum
	LiabilityGroup
	// RevenueGroup enum
	RevenueGroup
	// IncomeGroup enum
	IncomeGroup
	// ExpenseGroup enum
	ExpenseGroup
	// Branch
	// Warehouse
)

func (l LedgerGroupType) String() string {
	return [...]string{"Asset", "Liability", "Revenue", "Income", "Expense"}[l]
}
