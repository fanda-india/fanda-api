package scopes

import (
	"fanda-api/models"
	"fanda-api/options"

	"gorm.io/gorm"
)

// Paginate scope
func Paginate(opts options.ListOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		switch {
		case opts.Size > 100:
			opts.Size = 100
		case opts.Size <= 0 && opts.Page > 0:
			opts.Size = 10
		case opts.Size <= 0:
			opts.Size = 0
		}
		if opts.Page == 0 {
			opts.Page = 1
		}

		offset := (opts.Page - 1) * opts.Size

		// skip pagination
		if offset == 0 && opts.Size == 0 {
			return db
		}
		return db.Offset(offset).Limit(opts.Size)
	}
}

// All scope
func All(opts options.ListOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if !opts.All {
			return db.Where("active = ?", true)
		}
		return db
	}
}

// OrgID scope
func OrgID(orgID models.OrgID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("org_id = ?", orgID)
	}
}
