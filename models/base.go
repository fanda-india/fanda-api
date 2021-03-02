package models

import (
	"time"
)

// Base model
type Base struct {
	ID uint `gorm:"primarykey"`
}

// Audit model
type Audit struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
