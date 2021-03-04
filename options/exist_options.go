package options

import (
	"fanda-api/enums"
	"fanda-api/models"
)

// ExistOptions type
type ExistOptions struct {
	Field    enums.KeyField
	Value    string
	ParentID models.ID
}
