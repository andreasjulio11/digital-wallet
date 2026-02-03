package services

import (
	"digital-wallet/internal/config"
	"digital-wallet/internal/models"
	"digital-wallet/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.UserRepository.FindByEmail(email)

	if err != nil {
		return "", errors.New("Periksa kembali email anda")
	}

	if user.State != 1 {
		return "", errors.New("Pengguna Tidak Aktif")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Email Atau Password Salah")
	}

	token, _ := config.GenerateToken(int(user.ID))

	return token, nil
}

func (s *UserService) Register(name, email, password string) (models.User, error) {
	_, err := s.UserRepository.FindByEmail(email)

	if err == nil {
		return models.User{}, errors.New("Email sudah terdaftar")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, err
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	newUser := models.User{
		Name: name, Email: email, Password: string(passwordHash), Balanced: 0, State: 1,
	}

	err = s.UserRepository.Create(newUser)
	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

func (s *UserService) Profile(userID int) (models.User, error) {
	user, err := s.UserRepository.FindById(userID)
	if err != nil {
		return models.User{}, errors.New("Gagal Memproses")
	}
	return user, nil
}
