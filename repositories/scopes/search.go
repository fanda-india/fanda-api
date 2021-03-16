package scopes

import (
	"database/sql"
	"fanda-api/options"

	"gorm.io/gorm"
)

// SearchUser scope
func SearchUser(opts options.ListOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if opts.Search == "" {
			return db
		}
		opts.Search = "%" + opts.Search + "%"
		return db.Where("(user_name LIKE @search OR email LIKE @search OR "+
			"mobile_number LIKE @search OR first_name LIKE @search OR last_name LIKE @search)",
			sql.Named("search", opts.Search))
	}
}

// SearchDefault scope
func SearchDefault(opts options.ListOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if opts.Search == "" {
			return db
		}
		opts.Search = "%" + opts.Search + "%"
		return db.Where("(code LIKE @search OR name LIKE @search OR description LIKE @search)",
			sql.Named("search", opts.Search))
	}
}
