package options

import "fanda-api/models"

// ValidateOptions type
type ValidateOptions struct {
	ID     models.ID
	Code   string
	Name   string
	Email  string
	Mobile string
	Number string
	OrgID  models.OrgID
}
