package repository

import (
	"digital-wallet/internal/models"

	"gorm.io/gorm"
)

type DetailTransactionRepository struct {
	DB *gorm.DB
}

func (r *DetailTransactionRepository) GetDB() *gorm.DB {
	return r.DB
}
func (r *DetailTransactionRepository) Create(tx *gorm.DB, detailTransaction models.DetailTransaction) error {
	if tx != nil {
		return tx.Create(&detailTransaction).Error
	}
	return r.DB.Create(&detailTransaction).Error
}

func (r *DetailTransactionRepository) FindByUSerId(id int) ([]models.DetailTransaction, error) {
	var dt []models.DetailTransaction
	err := r.DB.Where("user_id = ?", id).Order("created_at desc").Find(&dt).Error
	return dt, err
}
