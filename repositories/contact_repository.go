package repositories

import (
	"fanda-api/models"

	"gorm.io/gorm"
)

// ContactRepository service
type ContactRepository struct{}

// NewContactRepository method
//func NewContactRepository() *ContactRepository {
//	return &ContactRepository{}
//}

// Save method
func (repo *ContactRepository) Save(contact *models.Contact, tx *gorm.DB) (*models.ID, error) {
	if contact.ID == 0 {
		// create
		if err := tx.Create(&contact).Error; err != nil {
			return nil, err
		}
		return &contact.ID, nil
	}
	// delete
	if contact.IsEmpty() {
		return repo.Delete(contact.ID, tx)
	}
	// update
	if err := tx.Save(&contact).Error; err != nil {
		return nil, err
	}
	return &contact.ID, nil
}

// Delete method
func (repo *ContactRepository) Delete(id models.ID, tx *gorm.DB) (*models.ID, error) {
	if err := tx.Delete(&models.User{}, id).Error; err != nil {
		return &id, err
	}
	return nil, nil
}
