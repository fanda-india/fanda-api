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
	var banks []models.Bank

	if err := repo.db.
		Model(&models.Bank{}).
		Joins("Ledger").
		Scopes(scopes.Paginate(opts), scopes.All(opts), scopes.SearchDefault(opts)).
		Find(&banks).Error; err != nil {
		return nil, err
	}

	// convert bank to bankdto
	banksDto := make([]dtos.BankDto, len(banks))
	for i, v := range banks {
		banksDto[i].FromBank(v)

	}
	count, err := repo.GetCount(opts)
	if err != nil {
		return nil, err
	}
	return &options.ListResult{Data: &banksDto, Count: count}, nil
}

// Read method
func (repo *BankRepository) Read(id models.ID) (*dtos.BankDto, error) {
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
	return (&dtos.BankDto{}).FromBank(bank), nil
}

// Create method
func (repo *BankRepository) Create(orgID models.ID, dto *dtos.BankDto) error {
	// validate
	var opts = options.ValidateOptions{
		ID: dto.ID, Code: dto.Code, Name: dto.Name,
		Number: *dto.AccountNumber, ParentID: orgID,
	}
	_, err := repo.Validate(opts)
	if err != nil {
		return err
	}

	// create
	groupID, err := getLedgerGroupID(repo.db, "BANKS")
	if err != nil {
		return err
	}
	bank := dto.ToBank()
	bank.Ledger.GroupID = groupID
	bank.Ledger.LedgerType = enums.Bank
	bank.Ledger.OrgID = orgID
	if err := repo.db.Create(&bank).Error; err != nil {
		return err
	}
	dto.ID = bank.ID
	return nil
}

// Update method
func (repo *BankRepository) Update(orgID models.ID, id models.ID, bank *models.Bank) error {
	var existOpts = options.ExistOptions{ID: id, Field: enums.IDField}
	if existID, err := repo.CheckExists(existOpts); err != nil {
		return err
	} else if id != existID {
		return options.NewNotFoundError("Bank")
	}
	bank.ID = id

	// validate
	var opts = options.ValidateOptions{
		ID: bank.ID, Code: bank.Ledger.Code, Name: bank.Ledger.Name,
		Number: *bank.AccountNumber, ParentID: orgID,
	}
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
	// condition := models.Bank{Ledger: &models.Ledger{}}
	var id models.ID
	var err error

	switch opts.Field {
	case enums.IDField:
		// condition.ID = opts.ID
		err = repo.db.Model(&models.Bank{}).
			Select("id").
			Where("id = ?", opts.ID).
			Scan(&id).Error
	case enums.CodeField:
		// condition.Ledger.Code = opts.Value
		condition := models.Ledger{Code: opts.Value, OrgID: opts.ParentID}
		err = repo.db.Model(&models.Ledger{}).
			Select("id").
			Where(condition).
			Scan(&id).Error
	case enums.NameField:
		// condition.Ledger.Name = opts.Value
		condition := models.Ledger{Name: opts.Value, OrgID: opts.ParentID}
		err = repo.db.Model(&models.Ledger{}).
			Select("id").
			Where(condition).
			Scan(&id).Error
	case enums.NumberField:
		// condition.AccountNumber = &opts.Value
		condition := models.Bank{AccountNumber: &opts.Value}
		err = repo.db.Model(&models.Bank{}).
			Select("id").
			Where(condition).
			Scan(&id).Error
	default:
		return 0, fmt.Errorf("CheckExists - Unknown field: %d", opts.Field)
	}
	if err != nil {
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
	exOpt := options.ExistOptions{Field: enums.CodeField, Value: opts.Code, ParentID: opts.ParentID}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != opts.ID {
		return false, errors.New("bank code already exists")
	}
	// Bank Name
	exOpt = options.ExistOptions{Field: enums.NameField, Value: opts.Name, ParentID: opts.ParentID}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != opts.ID {
		return false, errors.New("bank name already exists")
	}
	// Account number
	exOpt = options.ExistOptions{Field: enums.NumberField, Value: opts.Number, ParentID: opts.ParentID}
	if id, err := repo.CheckExists(exOpt); err != nil {
		return false, err
	} else if id != 0 && id != opts.ID {
		return false, errors.New("account number already exists")
	}
	return true, nil
}
