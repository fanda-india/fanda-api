package controllers

import "gorm.io/gorm"

var db *gorm.DB

// InitDatabase method
func InitDatabase(d *gorm.DB) {
	db = d
}
