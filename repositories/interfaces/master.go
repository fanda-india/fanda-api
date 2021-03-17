package interfaces

import (
	"fanda-api/models"
	"fanda-api/options"
)

type MasterRepository interface {
	List(options.ListOptions) (*options.ListResult, error)
	Read(id models.ID)
}
