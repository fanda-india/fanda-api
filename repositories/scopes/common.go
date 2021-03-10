package scopes

import (
	"fanda-api/models"
	"fanda-api/options"

	"gorm.io/gorm"
)

// Paginate scope
func Paginate(o options.ListOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		switch {
		case o.Size > 100:
			o.Size = 100
		case o.Size <= 0 && o.Page > 0:
			o.Size = 10
		case o.Size <= 0:
			o.Size = 0
		}
		if o.Page == 0 {
			o.Page = 1
		}

		offset := (o.Page - 1) * o.Size

		// skip pagination
		if offset == 0 && o.Size == 0 {
			return db
		}
		return db.Offset(offset).Limit(o.Size)
	}
}

// All scope
func All(o options.ListOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if !o.All {
			return db.Where("active = ?", true)
		}
		return db
	}
}

// OrgID scope
func OrgID(orgID models.ID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("org_id = ?", orgID)
	}
}
