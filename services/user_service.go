package services

import (
	"database/sql"
	"errors"
	"fanda-api/controllers/scopes"
	"fanda-api/dtos"
	"fanda-api/enums"
	"fanda-api/models"
	"fanda-api/options"
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
			return nil, errors.New("User not found")
		default:
			return nil, err
		}
	}
	return &user, nil
}

// Create method
func (s *UserService) Create(userDto dtos.UserDto) (*dtos.UserDto, error) {
	// var user models.User
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&user); err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	// 	return
	// }
	// defer r.Body.Close()

	var user = userDto.ToUser()
	if err := s.db.Create(&user).Error; err != nil {
		// respondWithError(w, http.StatusInternalServerError, err.Error())
		return nil, err
	}

	// apiuser := apiUser{ID: user.ID, UserName: user.UserName, Email: user.Email, MobileNumber: user.MobileNumber,
	// 	FirstName: user.FirstName, LastName: user.LastName, Active: user.Active}
	// w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
	// respondWithJSON(w, http.StatusCreated, apiuser)
	return userDto.FromUser(user), nil
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
	// condition := make(map[string]interface{})
	condition := models.User{}

	switch option.Field {
	case enums.Name:
		//condition["user_name"] = o.Value
		condition.UserName = option.Value
	case enums.Email:
		//condition["email"] = o.Value
		condition.Email = option.Value
	case enums.Mobile:
		// condition["mobile_number"] = o.Value
		condition.MobileNumber = option.Value
	default:
		return 0, fmt.Errorf("checkExists - Unknown field: %d", option.Field)
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
		return false, errors.New("Username is required") //&options.ValidateResult{Success: false, Error: "Username is required"}
	}
	if option.Email == "" {
		return false, errors.New("Email is required") //&options.ValidateResult{Success: false, Error: "Email is required"}
	}
	if option.Mobile == "" {
		return false, errors.New("Mobile number is required") //&options.ValidateResult{Success: false, Error: "Mobile number is required"}
	}

	// Duplicate validations
	// Username
	existOption := options.ExistOptions{Field: enums.Code, Value: option.Name}
	if id, err := s.CheckExists(existOption); err != nil {
		return false, err
	} else if id != uint(option.ID) {
		return false, errors.New("Username already exists")
	}
	// Email
	existOption = options.ExistOptions{Field: enums.Email, Value: option.Email}
	if id, err := s.CheckExists(existOption); err != nil {
		return false, err
	} else if id != uint(option.ID) {
		return false, errors.New("Email already exists")
	}
	// Mobile number
	existOption = options.ExistOptions{Field: enums.Mobile, Value: option.Mobile}
	if id, err := s.CheckExists(existOption); err != nil {
		return false, err
	} else if id != uint(option.ID) {
		return false, errors.New("Mobile number already exists")
	}
	return true, nil
}
