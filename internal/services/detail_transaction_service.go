package services

import (
	"digital-wallet/internal/models"
	"digital-wallet/internal/repository"
	"errors"
)

type DetailTransactionService struct {
	DetailTransactionRepository *repository.DetailTransactionRepository
	UserRepository              *repository.UserRepository
}

func (s *DetailTransactionService) CreateTransaction(userID int, amount float64, typeTranscation string) (models.DetailTransaction, error) {
	if amount < 0 || amount == 0 || amount == -0 {
		return models.DetailTransaction{}, errors.New("Pastikan kembali saldo anda")
	}

	transaction := s.DetailTransactionRepository.GetDB().Begin()

	user, err := s.UserRepository.FindById(userID)

	if err != nil {
		transaction.Rollback()
		return models.DetailTransaction{}, errors.New("User Tidak Ditemukan")
	}

	switch typeTranscation {
	case "topup":
		user.Balanced += amount
	case "withdraw":
		if user.Balanced < amount {
			transaction.Rollback()
			return models.DetailTransaction{}, errors.New("Saldo Tidak Cukup")
		}
		user.Balanced -= amount
	default:
		transaction.Rollback()
		return models.DetailTransaction{}, errors.New("tipe transaksi tidak dikenal")
	}

	if err := s.UserRepository.UpdateSaldo(transaction, user, user.Balanced); err != nil {
		transaction.Rollback()
		return models.DetailTransaction{}, err
	}

	newTransaction := models.DetailTransaction{
		UserID:          userID,
		Amount:          amount,
		TransactionType: typeTranscation,
	}
	if err := s.DetailTransactionRepository.Create(transaction, newTransaction); err != nil {
		transaction.Rollback()
		return models.DetailTransaction{}, errors.New("Gagal Menyimpan Riwayat Transaksi")
	}
	transaction.Commit()

	return newTransaction, nil
}

func (s *DetailTransactionService) SearchTransaction(userID int) ([]models.DetailTransaction, error) {
	transaction, err := s.DetailTransactionRepository.FindByUSerId(userID)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
