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
		Scopes(scopes.Paginate(opts), scopes.All(opts), scopes.SearchOrg(opts)).
		Find(&orgs).Error; err != nil {
		return nil, err
	}
	count, err := repo.GetCount(opts)
	if err != nil {
		return nil, err
	}
	return &options.ListResult{Data: &orgs, Count: count}, nil

}

// Read method
func (repo *OrganizationRepository) Read(id models.ID) (*models.Organization, error) {
	var org models.Organization

	if err := repo.db.First(&org, id).Error; err != nil {
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
func (repo *OrganizationRepository) Create(org *models.Organization) (*models.Organization, error) {
	// validate
	var opts = options.ValidateOptions{ID: org.ID, Code: org.Code, Name: org.Name}
	_, err := repo.Validate(opts)
	if err != nil {
		return nil, err
	}

	// create
	if err := repo.db.Create(&org).Error; err != nil {
		return nil, err
	}
	return org, nil
}

// Update method
func (repo *OrganizationRepository) Update(id models.ID, org *models.Organization) (*models.Organization, error) {
	// check record exists
	var exists bool
	if err := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM organizations WHERE id = ?)", id).Scan(&exists).Error; err != nil {
		return nil, err
	}
	if !exists {
		return nil, options.NewNotFoundError("Organization")
	}
	org.ID = id

	// validate
	var opts = options.ValidateOptions{ID: org.ID, Code: org.Code, Name: org.Name}
	_, err := repo.Validate(opts)
	if err != nil {
		return nil, err
	}

	// update
	if err := repo.db.Model(&models.Organization{}).
		Where("id = ?", id).
		Updates(org).Error; err != nil {
		return nil, err
	}
	return org, nil
}

// Delete method
func (repo *OrganizationRepository) Delete(id models.ID) (bool, error) {
	// check record exists
	var exists bool
	if err := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM organizations WHERE id = ?)", id).Scan(&exists).Error; err != nil {
		return false, err
	}
	if !exists {
		return false, options.NewNotFoundError("Organization")
	}

	if err := repo.db.Delete(&models.Organization{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetCount method
func (repo *OrganizationRepository) GetCount(opts options.ListOptions) (int64, error) {
	var count int64
	if err := repo.db.Model(&models.Organization{}).
		Scopes(scopes.All(opts), scopes.SearchOrg(opts)).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

// CheckExists method
func (repo *OrganizationRepository) CheckExists(opts options.ExistOptions) (models.ID, error) {
	condition := models.Organization{}

	switch opts.Field {
	case enums.Code:
		condition.Code = opts.Value
	case enums.Name:
		condition.Name = opts.Value
	default:
		return 0, fmt.Errorf("CheckExists - Unknown field: %d", opts.Field)
	}

	var id models.ID
	if err := repo.db.Model(&models.Organization{}).Select("id").Where(&condition).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

// Validate method
func (repo *OrganizationRepository) Validate(opts options.ValidateOptions) (bool, error) {
	// Required validations
	if opts.Code == "" {
		return false, errors.New("Org. code is required")
	}
	if opts.Name == "" {
		return false, errors.New("Org. name is required")
	}

	// Duplicate validations
	// Code
	exOpt := options.ExistOptions{Field: enums.Code, Value: opts.Code}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != uint(opts.ID) {
		return false, errors.New("Org. code already exists")
	}
	// Name
	exOpt = options.ExistOptions{Field: enums.Name, Value: opts.Name}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != uint(opts.ID) {
		return false, errors.New("Org. name already exists")
	}
	return true, nil
}
