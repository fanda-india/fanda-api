package models

import "gorm.io/gorm"

// DBContext type
type DBContext struct {
	DB *gorm.DB
}

// NewDBContext method
func NewDBContext() *DBContext {
	return &DBContext{}
}

// Initialize method
func (dbc *DBContext) Initialize(db *gorm.DB) {
	dbc.DB = db
}

// Migrate method
func (dbc *DBContext) Migrate() {
	dbc.DB.AutoMigrate(&User{})
	dbc.DB.AutoMigrate(&Address{})
	dbc.DB.AutoMigrate(&Contact{})
	dbc.DB.AutoMigrate(&Organization{})
	dbc.DB.AutoMigrate(&LedgerGroup{})
}
