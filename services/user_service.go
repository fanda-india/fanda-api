package services

import (
	"database/sql"
	"errors"
	"fanda-api/dtos"
	"fanda-api/enums"
	"fanda-api/models"
	"fanda-api/options"
	"fanda-api/services/scopes"
	"fmt"

	"gorm.io/gorm"
)

// UserService service
type UserService struct {
	db *gorm.DB
}

// NewUserService method
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

// List method
func (s *UserService) List(o options.ListOptions) (*options.ListResult, error) {
	var users []dtos.UserDto

	if err := s.db.Model(&models.User{}).
		Scopes(scopes.Paginate(o), scopes.All(o), scopes.SearchUser(o)).
		Find(&users).Error; err != nil {
		return nil, err
	}
	count, err := s.GetCount(o)
	if err != nil {
		return nil, err
	}
	return &options.ListResult{Data: &users, Count: count}, nil

}

// Read method
func (s *UserService) Read(id models.ID) (*dtos.UserDto, error) {
	var user dtos.UserDto

	if err := s.db.Model(&models.User{}).First(&user, id).Error; err != nil {
		switch err {
		case sql.ErrNoRows:
		case gorm.ErrRecordNotFound:
			return nil, options.NewNotFoundError("User")
		default:
			return nil, err
		}
	}
	return &user, nil
}

// Create method
func (s *UserService) Create(userDto *dtos.UserDto) (*dtos.UserDto, error) {
	var user = userDto.ToUser()

	// validate
	var vo = options.ValidateOptions{ID: user.ID, Name: user.UserName, Email: user.Email, Mobile: user.MobileNumber}
	_, err := s.Validate(vo)
	if err != nil {
		return nil, err
	}

	// create
	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return userDto.FromUser(user), nil
}

// Update method
func (s *UserService) Update(id models.ID, userDto *dtos.UserDto) (*dtos.UserDto, error) {
	// check record exists
	var exists bool
	if err := s.db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists).Error; err != nil {
		return nil, err
	}
	if !exists {
		return nil, options.NewNotFoundError("User")
	}

	var user = userDto.ToUser()

	// validate
	var vo = options.ValidateOptions{ID: user.ID, Name: user.UserName, Email: user.Email, Mobile: user.MobileNumber}
	_, err := s.Validate(vo)
	if err != nil {
		return nil, err
	}

	// update
	user.ID = 0
	if err := s.db.Model(&models.User{}).
		Where("id = ?", id).
		Updates(user).Error; err != nil {
		return nil, err
	}
	return userDto.FromUser(user), nil
}

// Delete method
func (s *UserService) Delete(id models.ID) (bool, error) {
	// check record exists
	var exists bool
	if err := s.db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists).Error; err != nil {
		return false, err
	}
	if !exists {
		return false, options.NewNotFoundError("User")
	}

	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetCount method
func (s *UserService) GetCount(o options.ListOptions) (int64, error) {
	var count int64
	if err := s.db.Model(&models.User{}).
		Scopes(scopes.All(o), scopes.SearchUser(o)).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

// CheckExists method
func (s *UserService) CheckExists(option options.ExistOptions) (models.ID, error) {
	condition := models.User{}

	switch option.Field {
	case enums.Name:
		condition.UserName = option.Value
	case enums.Email:
		condition.Email = option.Value
	case enums.Mobile:
		condition.MobileNumber = option.Value
	default:
		return 0, fmt.Errorf("CheckExists - Unknown field: %d", option.Field)
	}

	var id models.ID
	if err := s.db.Model(&models.User{}).Select("id").Where(&condition).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

// Validate method
func (s *UserService) Validate(option options.ValidateOptions) (bool, error) {
	// Required validations
	if option.Name == "" {
		return false, errors.New("Username is required")
	}
	if option.Email == "" {
		return false, errors.New("Email is required")
	}
	if option.Mobile == "" {
		return false, errors.New("Mobile number is required")
	}

	// Duplicate validations
	// Username
	existOption := options.ExistOptions{Field: enums.Name, Value: option.Name}
	if id, err := s.CheckExists(existOption); err != nil {
		return false, err
	} else if id != 0 && id != uint(option.ID) {
		return false, errors.New("Username already exists")
	}
	// Email
	existOption = options.ExistOptions{Field: enums.Email, Value: option.Email}
	if id, err := s.CheckExists(existOption); err != nil {
		return false, err
	} else if id != 0 && id != uint(option.ID) {
		return false, errors.New("Email already exists")
	}
	// Mobile number
	existOption = options.ExistOptions{Field: enums.Mobile, Value: option.Mobile}
	if id, err := s.CheckExists(existOption); err != nil {
		return false, err
	} else if id != 0 && id != uint(option.ID) {
		return false, errors.New("Mobile number already exists")
	}
	return true, nil
}
