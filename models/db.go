package models

import "gorm.io/gorm"

// Migrate method
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Contact{})
	db.AutoMigrate(&Organization{})
	db.AutoMigrate(&LedgerGroup{})
}
