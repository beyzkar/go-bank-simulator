package services

import (
	"errors"

	"github.com/beyza/go-bank-simulator/models"
	"github.com/beyza/go-bank-simulator/repositorys"
)

func GetTransactionByID(id uint) (*models.Transaction, error) {
	if id == 0 {
		return nil, errors.New("geçersiz transaction id")
	}
	return repositorys.GetTransactionByID(id)
}

func GetTransactionsByAccountID(accountID uint) ([]models.Transaction, error) {
	if accountID == 0 {
		return nil, errors.New("geçersiz account id")
	}
	return repositorys.GetTransactionsByAccountID(accountID)
}

// ---- Business Logic ----

func Deposit(accountID uint, amount float64) (*models.Transaction, error) {
	if accountID == 0 {
		return nil, errors.New("geçersiz account id")
	}
	if amount <= 0 {
		return nil, errors.New("amount 0'dan büyük olmalı")
	}

	// DB işlemi repository'de
	return repositorys.Deposit(accountID, amount)
}

func Withdraw(accountID uint, amount float64) (*models.Transaction, error) {
	if accountID == 0 {
		return nil, errors.New("geçersiz account id")
	}
	if amount <= 0 {
		return nil, errors.New("amount 0'dan büyük olmalı")
	}

	// Yetersiz bakiye kontrolü repository içinde (transaction güvenliği için)
	return repositorys.Withdraw(accountID, amount)
}
