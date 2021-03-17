package repositories

import (
	"database/sql"
	"errors"
	"fanda-api/enums"
	"fanda-api/models"
	"fanda-api/options"
	"fanda-api/repositories/scopes"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// OrganizationRepository service
type OrganizationRepository struct {
	db *gorm.DB
}

// NewOrganizationRepository method
func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db}
}

// List method
func (repo *OrganizationRepository) List(opts options.ListOptions) (*options.ListResult, error) {
	var orgs []models.Organization

	if err := repo.db.
		Scopes(scopes.Paginate(opts), scopes.All(opts), scopes.SearchDefault(opts)).
		Find(&orgs).Error; err != nil {
		return nil, err
	}
	count, err := repo.Count(opts)
	if err != nil {
		return nil, err
	}
	return &options.ListResult{Data: &orgs, Count: count}, nil

}

// Read method
func (repo *OrganizationRepository) Read(id models.ID) (*models.Organization, error) {
	var org models.Organization

	if err := repo.db.
		Preload("Address").Preload("Contact").
		First(&org, id).Error; err != nil {
		switch err {
		case sql.ErrNoRows:
		case gorm.ErrRecordNotFound:
			return nil, options.NewNotFoundError("Organization")
		default:
			return nil, err
		}
	}
	return &org, nil
}

// Create method
func (repo *OrganizationRepository) Create(org *models.Organization) error {
	// validate
	var opts = options.ValidateOptions{ID: org.ID, Code: org.Code, Name: org.Name}
	_, err := repo.Validate(opts)
	if err != nil {
		return err
	}

	// create
	if err := repo.db.Create(&org).Error; err != nil {
		return err
	}
	return nil
}

// Update method
func (repo *OrganizationRepository) Update(id models.ID, org *models.Organization) error {
	// check record exists
	// var exists bool
	// if err := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM organizations WHERE id = ?)", id).Scan(&exists).Error; err != nil {
	// 	return nil, err
	// }
	var existOpts = options.ExistOptions{ID: id, Field: enums.IDField}
	if existID, err := repo.Exists(existOpts); err != nil {
		return err
	} else if id != existID {
		return options.NewNotFoundError("Organization")
	}
	org.ID = id

	// validate
	var opts = options.ValidateOptions{ID: org.ID, Code: org.Code, Name: org.Name}
	_, err := repo.Validate(opts)
	if err != nil {
		return err
	}

	// update
	// if err := repo.db.Model(&models.Organization{}).
	// 	Where("id = ?", id).
	// 	Updates(org).Error; err != nil {
	// 	return nil, err
	// }
	if err := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(&org).Error; err != nil {
		return err
	}
	return nil
}

// Delete method
func (repo *OrganizationRepository) Delete(id models.ID) (bool, error) {
	// check record exists
	// var exists bool
	// if err := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM organizations WHERE id = ?)", id).Scan(&exists).Error; err != nil {
	// 	return false, err
	// }
	var opts = options.ExistOptions{ID: id, Field: enums.IDField}
	if existID, err := repo.Exists(opts); err != nil {
		return false, err
	} else if id != existID {
		return false, options.NewNotFoundError("Organization")
	}

	org, err := repo.Read(id)
	if err != nil {
		return false, err
	}
	if err := repo.db.Select(clause.Associations).Delete(&org).Error; err != nil {
		return false, err
	}
	if err := repo.db.Delete(&org.Address).Error; err != nil {
		return false, err
	}
	if err := repo.db.Delete(&org.Contact).Error; err != nil {
		return false, err
	}
	return true, nil
}

// Count method
func (repo *OrganizationRepository) Count(opts options.ListOptions) (int64, error) {
	var count int64
	if err := repo.db.Model(&models.Organization{}).
		Scopes(scopes.All(opts), scopes.SearchDefault(opts)).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

// Exists method
func (repo *OrganizationRepository) Exists(opts options.ExistOptions) (models.ID, error) {
	if opts.Value == "" && opts.ID == 0 {
		return 0, errors.New("value is required")
	}

	var id uint
	var err error
	db := repo.db.Model(&models.Organization{}).Select("id")

	switch opts.Field {
	case enums.IDField:
		// condition.ID = opts.ID
		err = db.Where("id = ?", opts.ID).Scan(&id).Error
	case enums.CodeField:
		// condition.Code = opts.Value
		err = db.Where("code = ?", opts.Value).Scan(&id).Error
	case enums.NameField:
		// condition.Name = opts.Value
		err = db.Where("name = ?", opts.Value).Scan(&id).Error
	default:
		return 0, fmt.Errorf("Exists - Unknown field: %d", opts.Field)
	}

	if err != nil {
		return 0, err
	}
	return models.ID(id), nil
}

// Validate method
func (repo *OrganizationRepository) Validate(opts options.ValidateOptions) (bool, error) {
	// Required validations
	if opts.Code == "" {
		return false, errors.New("org. code is required")
	}
	if opts.Name == "" {
		return false, errors.New("org. name is required")
	}

	// Duplicate validations
	// Code
	exOpt := options.ExistOptions{Field: enums.CodeField, Value: opts.Code}
	if id, err := repo.Exists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != opts.ID {
		return false, errors.New("org. code already exists")
	}
	// Name
	exOpt = options.ExistOptions{Field: enums.NameField, Value: opts.Name}
	if id, err := repo.Exists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != opts.ID {
		return false, errors.New("org. name already exists")
	}
	return true, nil
}
