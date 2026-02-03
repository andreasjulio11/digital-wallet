package repository

import (
	"digital-wallet/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *UserRepository) Create(user models.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepository) FindById(userID int) (models.User, error) {
	var user models.User
	err := r.DB.Where("id = ?", userID).First(&user).Error
	return user, err
}

func (r *UserRepository) UpdateSaldo(tx *gorm.DB, user models.User, newBalance float64) error {

	return tx.Model(&user).Update("balanced", newBalance).Error
}
