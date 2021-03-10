package dtos

import "fanda-api/models"

// UserDto model
type UserDto struct {
	ID           models.ID `json:"id,omitempty"`
	UserName     string    `json:"userName,omitempty"`
	Email        string    `json:"email,omitempty"`
	MobileNumber string    `json:"mobileNumber,omitempty"`
	Password     string    `json:"password,omitempty"`
	FirstName    *string   `json:"firstName,omitempty"`
	LastName     *string   `json:"lastName,omitempty"`
	Active       *bool     `json:"active"`
}

// ToUser method
func (u *UserDto) ToUser() *models.User {
	return &models.User{ID: u.ID, UserName: u.UserName,
		Email: u.Email, MobileNumber: u.MobileNumber,
		Password: u.Password, FirstName: u.FirstName,
		LastName: u.LastName, Active: u.Active,
	}
}

// FromUser method
func (u *UserDto) FromUser(user *models.User) *UserDto {
	u.ID = user.ID
	u.UserName = user.UserName
	u.Email = user.Email
	u.MobileNumber = user.MobileNumber
	u.Password = user.Password
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Active = user.Active
	return u
}
