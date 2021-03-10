package repositories

import (
	"database/sql"
	"errors"
	"fanda-api/dtos"
	"fanda-api/enums"
	"fanda-api/models"
	"fanda-api/options"
	"fanda-api/repositories/scopes"
	"fmt"

	"gorm.io/gorm"
)

// UserRepository service
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository method
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// List method
func (repo *UserRepository) List(opts options.ListOptions) (*options.ListResult, error) {
	var users []dtos.UserDto

	if err := repo.db.Model(&models.User{}).
		Scopes(scopes.Paginate(opts), scopes.All(opts), scopes.SearchUser(opts)).
		Find(&users).Error; err != nil {
		return nil, err
	}
	count, err := repo.GetCount(opts)
	if err != nil {
		return nil, err
	}
	return &options.ListResult{Data: &users, Count: count}, nil

}

// Read method
func (repo *UserRepository) Read(id models.ID) (*dtos.UserDto, error) {
	var user dtos.UserDto

	if err := repo.db.Model(&models.User{}).First(&user, id).Error; err != nil {
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
func (repo *UserRepository) Create(userDto *dtos.UserDto) error {
	var user = userDto.ToUser()
	// validate
	var opts = options.ValidateOptions{ID: user.ID, Name: user.UserName, Email: user.Email, Mobile: user.MobileNumber}
	_, err := repo.Validate(opts)
	if err != nil {
		return err
	}
	// create
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}
	userDto.FromUser(user)
	return nil
}

// Update method
func (repo *UserRepository) Update(id models.ID, userDto *dtos.UserDto) error {
	// check record exists
	// var exists bool
	// if err := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists).Error; err != nil {
	// 	return nil, err
	// }
	userDto.ID = id
	var existOpts = options.ExistOptions{ID: id, Field: enums.IDField}
	if existID, err := repo.CheckExists(existOpts); err != nil {
		return err
	} else if id != existID {
		return options.NewNotFoundError("User")
	}
	var user = userDto.ToUser()

	// validate
	var opts = options.ValidateOptions{ID: user.ID, Name: user.UserName, Email: user.Email, Mobile: user.MobileNumber}
	_, err := repo.Validate(opts)
	if err != nil {
		return err
	}

	// update
	// user.ID = 0
	if err := repo.db.
		Model(&models.User{}).
		//Select("UserName", "Email", "MobileNumber", "Password", "FirstName", "LastName", "Active").
		Omit("id").
		Where("id = ?", id).
		Updates(&user).Error; err != nil {
		return err
	}
	// user.ID = id
	userDto.FromUser(user)
	return nil
}

// Delete method
func (repo *UserRepository) Delete(id models.ID) (bool, error) {
	// check record exists
	// var exists bool
	// if err := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists).Error; err != nil {
	// 	return false, err
	// }
	var opts = options.ExistOptions{ID: id, Field: enums.IDField}
	if existID, err := repo.CheckExists(opts); err != nil {
		return false, err
	} else if id != existID {
		return false, options.NewNotFoundError("User")
	}

	if err := repo.db.Delete(&models.User{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetCount method
func (repo *UserRepository) GetCount(opts options.ListOptions) (int64, error) {
	var count int64
	if err := repo.db.Model(&models.User{}).
		Scopes(scopes.All(opts), scopes.SearchUser(opts)).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

// CheckExists method
func (repo *UserRepository) CheckExists(opts options.ExistOptions) (models.ID, error) {
	if opts.Value == "" && opts.ID == 0 {
		return 0, errors.New("Value is required")
	}

	condition := models.User{}

	switch opts.Field {
	case enums.IDField:
		condition.ID = opts.ID
	case enums.NameField:
		condition.UserName = opts.Value
	case enums.EmailField:
		condition.Email = opts.Value
	case enums.MobileField:
		condition.MobileNumber = opts.Value
	default:
		return 0, fmt.Errorf("CheckExists - Unknown field: %d", opts.Field)
	}

	var id models.ID
	if err := repo.db.Model(&models.User{}).
		Select("id").
		Where(&condition).
		Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

// Validate method
func (repo *UserRepository) Validate(opts options.ValidateOptions) (bool, error) {
	// Required validations
	if opts.Name == "" {
		return false, errors.New("Username is required")
	}
	if opts.Email == "" {
		return false, errors.New("Email is required")
	}
	if opts.Mobile == "" {
		return false, errors.New("Mobile number is required")
	}

	// Duplicate validations
	// Username
	exOpt := options.ExistOptions{Field: enums.NameField, Value: opts.Name}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != uint(opts.ID) {
		return false, errors.New("Username already exists")
	}
	// Email
	exOpt = options.ExistOptions{Field: enums.EmailField, Value: opts.Email}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != uint(opts.ID) {
		return false, errors.New("Email already exists")
	}
	// Mobile number
	exOpt = options.ExistOptions{Field: enums.MobileField, Value: opts.Mobile}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != uint(opts.ID) {
		return false, errors.New("Mobile number already exists")
	}
	return true, nil
}
