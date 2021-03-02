package models

// LedgerGroup model
type LedgerGroup struct {
	Base
	Code        string
	Name        string
	Description *string
	GroupType   int
	ParentID    *int
}
