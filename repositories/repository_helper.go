package repositories

import (
	"fanda-api/models"

	"gorm.io/gorm"
)

func getLedgerGroupID(tx *gorm.DB, groupCode string) (*models.ID, error) {
	var id uint

	if err := tx.Model(&models.LedgerGroup{}).
		Select("id").
		Where("code = ?", groupCode).
		Scan(&id).Error; err != nil {
		return nil, err
	}
	groupID := models.ID(id)
	return &groupID, nil
}
