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

// LedgerRepository service
type LedgerRepository struct {
	db *gorm.DB
}

// NewLedgerRepository method
func NewLedgerRepository(db *gorm.DB) *LedgerRepository {
	return &LedgerRepository{db}
}

// List method
func (repo *LedgerRepository) List(orgID models.OrgID, opts options.ListOptions) (*options.ListResult, error) {
	var ledgers []models.Ledger

	if err := repo.db.
		Scopes(scopes.Paginate(opts), scopes.OrgID(orgID), scopes.All(opts), scopes.SearchDefault(opts)).
		Find(&ledgers).Error; err != nil {
		return nil, err
	}
	count, err := repo.Count(orgID, opts)
	if err != nil {
		return nil, err
	}
	return &options.ListResult{Data: &ledgers, Count: count}, nil
}

// Read method
func (repo *LedgerRepository) Read(id models.ID, orgID models.OrgID) (*models.Ledger, error) {
	var ledger models.Ledger

	if err := repo.db.
		First(&ledger, map[string]interface{}{"id": id, "org_id": orgID}).Error; err != nil {
		switch err {
		case sql.ErrNoRows:
		case gorm.ErrRecordNotFound:
			return nil, options.NewNotFoundError("Ledger")
		default:
			return nil, err
		}
	}
	return &ledger, nil
}

// Create method
func (repo *LedgerRepository) Create(orgID models.OrgID, ledger *models.Ledger) error {
	// validate
	var opts = options.ValidateOptions{ID: ledger.ID, Name: ledger.Name, OrgID: orgID}
	_, err := repo.Validate(opts)
	if err != nil {
		return err
	}

	// create
	ledger.OrgID = orgID
	if err := repo.db.Create(&ledger).Error; err != nil {
		return err
	}
	return nil
}

// Update method
func (repo *LedgerRepository) Update(id models.ID, orgID models.OrgID, ledger *models.Ledger) error {
	// check record exists
	// var exists bool
	// if err := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM ledgers WHERE id = ?)", id).Scan(&exists).Error; err != nil {
	// 	return nil, err
	// }
	var existOpts = options.ExistOptions{ID: id, Field: enums.IDField}
	if existID, err := repo.Exists(existOpts); err != nil {
		return err
	} else if id != existID {
		return options.NewNotFoundError("Ledger")
	}
	ledger.ID = id

	// validate
	var opts = options.ValidateOptions{ID: ledger.ID, Name: ledger.Name, OrgID: orgID}
	_, err := repo.Validate(opts)
	if err != nil {
		return err
	}

	// update
	if err := repo.db.Model(&models.Ledger{}).
		Omit("id").
		Where("id = ?", id).
		Updates(&ledger).Error; err != nil {
		return err
	}
	// if err := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&ledger).Error; err != nil {
	// 	return nil, err
	// }
	return nil
}

// Delete method
func (repo *LedgerRepository) Delete(id models.ID) (bool, error) {
	// check record exists
	// var exists bool
	// if err := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM ledgers WHERE id = ?)", id).Scan(&exists).Error; err != nil {
	// 	return false, err
	// }
	var opts = options.ExistOptions{ID: id, Field: enums.IDField}
	if existID, err := repo.Exists(opts); err != nil {
		return false, err
	} else if id != existID {
		return false, options.NewNotFoundError("Ledger")
	}

	// ledger, err := repo.Read(orgId, id)
	// if err != nil {
	// 	return false, err
	// }
	if err := repo.db.
		//Select(clause.Associations).
		Delete(&models.Ledger{}, id).Error; err != nil {
		return false, err
	}
	// if err := repo.db.Delete(&ledger.Address).Error; err != nil {
	// 	return false, err
	// }
	// if err := repo.db.Delete(&ledger.Contact).Error; err != nil {
	// 	return false, err
	// }
	return true, nil
}

// Count method
func (repo *LedgerRepository) Count(orgID models.OrgID, opts options.ListOptions) (int64, error) {
	var count int64
	if err := repo.db.Model(&models.Ledger{}).
		Scopes(scopes.All(opts), scopes.SearchDefault(opts), scopes.OrgID(orgID)).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

// Exists method
func (repo *LedgerRepository) Exists(opts options.ExistOptions) (models.ID, error) {
	if opts.Value == "" && opts.ID == 0 {
		return 0, errors.New("value is required")
	}

	var id uint
	var err error
	db := repo.db.Model(&models.Ledger{}).Select("id")

	switch opts.Field {
	case enums.IDField:
		err = db.Where("id = ?", opts.ID).Scan(&id).Error
	case enums.NameField:
		err = db.Where("name = ? and org_id = ?", opts.Value, opts.OrgID).Scan(&id).Error
	default:
		return 0, fmt.Errorf("Exists - Unknown field: %d", opts.Field)
	}

	if err != nil {
		return 0, err
	}
	return models.ID(id), nil
}

// Validate method
func (repo *LedgerRepository) Validate(opts options.ValidateOptions) (bool, error) {
	// Required validations
	if opts.Code == "" {
		return false, errors.New("ledger code is required")
	}
	if opts.Name == "" {
		return false, errors.New("ledger name is required")
	}

	// Duplicate validations
	// Code
	// exOpt := options.ExistOptions{Field: enums.CodeField, Value: opts.Code, OrgID: opts.OrgID}
	// if id, err := repo.Exists(exOpt); err != nil {
	// 	return false, err
	// } else if id != 0 && id != opts.ID {
	// 	return false, errors.New("ledger code already exists")
	// }
	// Name
	exOpt := options.ExistOptions{Field: enums.NameField, Value: opts.Name, OrgID: opts.OrgID}
	if id, err := repo.Exists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != opts.ID {
		return false, errors.New("ledger name already exists")
	}
	return true, nil
}
