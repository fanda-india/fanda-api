package repositories

import (
	"fanda-api/models"

	"gorm.io/gorm"
)

func getLedgerGroupID(tx *gorm.DB, groupCode string) (*models.ID, error) {
	var groupID models.ID

	if err := tx.Model(&models.LedgerGroup{}).
		Select("id").
		Where("code = ?", groupCode).
		Scan(&groupID).Error; err != nil {
		return nil, err
	}
	return &groupID, nil
}
