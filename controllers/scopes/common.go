package scopes

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

// Paginate scope
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := r.URL.Query()
		page, _ := strconv.Atoi(query.Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(query.Get("size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

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
