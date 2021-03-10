package options

import (
	"fanda-api/enums"
	"fanda-api/models"
)

// ExistOptions type
type ExistOptions struct {
	ID       models.ID
	Field    enums.KeyField
	Value    string
	ParentID models.ID
}
