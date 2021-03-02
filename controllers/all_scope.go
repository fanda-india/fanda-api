package controllers

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

// All scope
func All(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := r.URL.Query()
		all, _ := strconv.ParseBool(query.Get("all"))

		if !all {
			return db.Where("active = ?", true)
		}
		return db
	}
}
