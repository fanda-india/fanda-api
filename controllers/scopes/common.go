package scopes

import (
	"fanda-api/options"

	"gorm.io/gorm"
)

// Paginate scope
func Paginate(o options.ListOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if o.Page == 0 {
			o.Page = 1
		}

		switch {
		case o.Size > 100:
			o.Size = 100
		case o.Size <= 0:
			o.Size = 10
		}

		offset := (o.Page - 1) * o.Size
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
