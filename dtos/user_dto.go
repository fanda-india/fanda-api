package dtos

import "fanda-api/models"

// UserDto type
type UserDto struct {
	ID           models.ID `json:"id"`
	UserName     string    `json:"userName"`
	Email        string    `json:"email"`
	MobileNumber string    `json:"mobileNumber"`
	FirstName    *string   `json:"firstName"`
	LastName     *string   `json:"lastName"`
	Active       bool      `json:"active"`
}
