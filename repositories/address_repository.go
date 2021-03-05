package repositories

import (
	"fanda-api/models"

	"gorm.io/gorm"
)

// AddressRepository service
type AddressRepository struct{}

// NewAddressRepository method
func NewAddressRepository() *AddressRepository {
	return &AddressRepository{}
}

// Save method
func (repo *AddressRepository) Save(address *models.Address, tx *gorm.DB) (*models.ID, error) {
	if address.ID == 0 {
		// create
		if err := tx.Create(&address).Error; err != nil {
			return nil, err
		}
		return &address.ID, nil
	}
	// delete
	if address.IsEmpty() {
		return repo.Delete(address.ID, tx)
	}
	// update
	if err := tx.Save(&address).Error; err != nil {
		return nil, err
	}
	return &address.ID, nil
}

// Delete method
func (repo *AddressRepository) Delete(id models.ID, tx *gorm.DB) (*models.ID, error) {
	if err := tx.Delete(&models.User{}, id).Error; err != nil {
		return &id, err
	}
	return nil, nil
}
