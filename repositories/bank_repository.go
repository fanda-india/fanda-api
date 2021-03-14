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

// BankRepository service
type BankRepository struct {
	db *gorm.DB
}

// NewBankRepository method
func NewBankRepository(db *gorm.DB) *BankRepository {
	return &BankRepository{db}
}

// List method
func (repo *BankRepository) List(orgID models.ID, opts options.ListOptions) (*options.ListResult, error) {
	var banks []models.Bank //dtos.BankDto

	// if err := repo.db.
	// 	Model(&models.Bank{}).
	// 	// Association("Ledger").
	// 	// Scopes(scopes.Paginate(opts), scopes.All(opts), scopes.SearchDefault(opts)).
	// 	Find(&banks).Error; err != nil {
	// 	return nil, err
	// }

	if err := repo.db.
		Model(&models.Bank{}).
		Joins("Ledger").
		Scopes(scopes.Paginate(opts), scopes.All(opts), scopes.SearchDefault(opts)).
		// Raw("select banks.id, ledgers.code, ledgers.name, ledgers.description," +
		// 	"banks.account_number, banks.account_type, banks.ifsc_code, banks.micr_code, banks.branch_code, banks.branch_name, banks.is_default," +
		// 	"ledgers.org_id, ledgers.active " +
		// 	"from banks inner join ledgers on banks.ledger_id=ledgers.id").
		Find(&banks).Error; err != nil {
		return nil, err
	}
	banksDto := make([]dtos.BankDto, len(banks))
	for i, v := range banks {
		banksDto[i].FromBank(&v)

	}
	// "where org_id = @orgId", sql.Named("orgId", orgID)
	count, err := repo.GetCount(opts)
	if err != nil {
		return nil, err
	}
	return &options.ListResult{Data: &banksDto, Count: count}, nil

}

// Read method
func (repo *BankRepository) Read(id models.ID) (*models.Bank, error) {
	var bank models.Bank

	if err := repo.db.
		Preload("Ledger").
		Preload("Address").Preload("Contact").
		First(&bank, id).Error; err != nil {
		switch err {
		case sql.ErrNoRows:
		case gorm.ErrRecordNotFound:
			return nil, options.NewNotFoundError("Bank")
		default:
			return nil, err
		}
	}
	return &bank, nil
}

// Create method
func (repo *BankRepository) Create(orgID models.ID, bank *models.Bank) error {
	// validate
	var opts = options.ValidateOptions{ID: bank.ID, Code: bank.Ledger.Code, Name: bank.Ledger.Name, Number: *bank.AccountNumber}
	_, err := repo.Validate(opts)
	if err != nil {
		return err
	}

	// create
	bank.Ledger.OrgID = orgID
	if err := repo.db.Create(&bank).Error; err != nil {
		return err
	}
	return nil
}

// Update method
func (repo *BankRepository) Update(id models.ID, bank *models.Bank) error {
	var existOpts = options.ExistOptions{ID: id, Field: enums.IDField}
	if existID, err := repo.CheckExists(existOpts); err != nil {
		return err
	} else if id != existID {
		return options.NewNotFoundError("Bank")
	}
	bank.ID = id

	// validate
	var opts = options.ValidateOptions{ID: bank.ID, Code: bank.Ledger.Code, Name: bank.Ledger.Name, Number: *bank.AccountNumber}
	_, err := repo.Validate(opts)
	if err != nil {
		return err
	}

	// update
	// if err := repo.db.Model(&models.Bank{}).
	// 	Omit("id").
	// 	Where("id = ?", id).
	// 	Updates(&bank).Error; err != nil {
	// 	return err
	// }
	if err := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(&bank).Error; err != nil {
		return err
	}
	return nil
}

// Delete method
func (repo *BankRepository) Delete(id models.ID) (bool, error) {
	var opts = options.ExistOptions{ID: id, Field: enums.IDField}
	if existID, err := repo.CheckExists(opts); err != nil {
		return false, err
	} else if id != existID {
		return false, options.NewNotFoundError("Bank")
	}

	bank, err := repo.Read(id)
	if err != nil {
		return false, err
	}
	if err := repo.db.
		//Select(clause.Associations).
		Delete(&models.Ledger{}, id).Error; err != nil {
		return false, err
	}
	if err := repo.db.Delete(&bank.Address).Error; err != nil {
		return false, err
	}
	if err := repo.db.Delete(&bank.Contact).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetCount method
func (repo *BankRepository) GetCount(opts options.ListOptions) (int64, error) {
	var count int64
	if err := repo.db.Model(&models.Bank{}).
		Joins("Ledger").
		Scopes(scopes.All(opts), scopes.SearchDefault(opts)).
		Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

// CheckExists method
func (repo *BankRepository) CheckExists(opts options.ExistOptions) (models.ID, error) {
	condition := models.Bank{}

	switch opts.Field {
	case enums.IDField:
		condition.ID = opts.ID
	case enums.CodeField:
		condition.Ledger.Code = opts.Value
	case enums.NameField:
		condition.Ledger.Name = opts.Value
	case enums.NumberField:
		condition.AccountNumber = &opts.Value
	default:
		return 0, fmt.Errorf("CheckExists - Unknown field: %d", opts.Field)
	}

	var id models.ID
	if err := repo.db.Model(&models.Bank{}).
		Preload("Ledger").
		Select("id").
		Where(&condition).
		Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

// Validate method
func (repo *BankRepository) Validate(opts options.ValidateOptions) (bool, error) {
	// Required validations
	if opts.Code == "" {
		return false, errors.New("bank code is required")
	}
	if opts.Name == "" {
		return false, errors.New("bank name is required")
	}
	if opts.Number == "" {
		return false, errors.New("account number is required")
	}

	// Duplicate validations
	// Bank Code
	exOpt := options.ExistOptions{Field: enums.CodeField, Value: opts.Code}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != opts.ID {
		return false, errors.New("bank code already exists")
	}
	// Bank Name
	exOpt = options.ExistOptions{Field: enums.NameField, Value: opts.Name}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != opts.ID {
		return false, errors.New("bank name already exists")
	}
	// Account number
	exOpt = options.ExistOptions{Field: enums.NumberField, Value: opts.Number}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != opts.ID {
		return false, errors.New("account number already exists")
	}
	return true, nil
}
