package controllers

import (
	"database/sql"
	"net/http"

	"gorm.io/gorm"
)

// SearchUser scope
func SearchUser(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := r.URL.Query()
		search := query.Get("search")

		if search == "" {
			return db
		}
		search = "%" + search + "%"
		return db.Where("(user_name LIKE @search OR email LIKE @search OR first_name LIKE @search OR last_name LIKE @search)",
			sql.Named("search", search))
	}
}
