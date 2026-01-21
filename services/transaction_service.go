package services

import (
	"errors"

	"github.com/beyza/go-bank-simulator/models"
	"github.com/beyza/go-bank-simulator/repositorys"
)

// =========================
// READ OPERATIONS
// =========================

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

func GetLastTransactionByAccountID(accountID uint) (*models.Transaction, error) {
	if accountID == 0 {
		return nil, errors.New("geçersiz account id")
	}
	return repositorys.FindLastTransactionByAccountID(accountID)
}

// =========================
// BUSINESS LOGIC
// =========================

func Deposit(accountID uint, amount float64) (*models.Transaction, error) {
	if accountID == 0 {
		return nil, errors.New("geçersiz account id")
	}
	if amount <= 0 {
		return nil, errors.New("amount 0'dan büyük olmalı")
	}
	return repositorys.Deposit(accountID, amount)
}

func Withdraw(accountID uint, amount float64) (*models.Transaction, error) {
	if accountID == 0 {
		return nil, errors.New("geçersiz account id")
	}
	if amount <= 0 {
		return nil, errors.New("amount 0'dan büyük olmalı")
	}
	return repositorys.Withdraw(accountID, amount)
}

// =========================
// TRANSFER (AccountID ile)
// =========================

func Transfer(fromAccountID, toAccountID uint, amount float64) (*models.Transaction, *models.Transaction, error) {
	if fromAccountID == 0 || toAccountID == 0 {
		return nil, nil, errors.New("geçersiz hesap id")
	}
	if fromAccountID == toAccountID {
		return nil, nil, errors.New("aynı hesaba transfer yapılamaz")
	}
	if amount <= 0 {
		return nil, nil, errors.New("tutar 0'dan büyük olmalı")
	}

	return repositorys.Transfer(fromAccountID, toAccountID, amount)
}

// =========================
// TRANSFER (CustomerID ile)
// - Gönderenin hesapları içinden EN YÜKSEK BAKİYELİ hesabı seçer
// - Alıcının ilk hesabına gönderir
//
// Not: "yetersiz bakiye" kontrolü repository.Transfer içinde yapılır
// çünkü güvenli kontrol DB transaction içinde olmalı.
// =========================

func TransferByCustomerID(fromCustomerID, toCustomerID uint, amount float64) (*models.Transaction, *models.Transaction, error) {
	if fromCustomerID == 0 || toCustomerID == 0 {
		return nil, nil, errors.New("geçersiz customer id")
	}
	if fromCustomerID == toCustomerID {
		return nil, nil, errors.New("aynı müşteriye transfer yapılamaz")
	}
	if amount <= 0 {
		return nil, nil, errors.New("tutar 0'dan büyük olmalı")
	}

	fromAccs, err := repositorys.GetAccountsByCustomerID(fromCustomerID)
	if err != nil {
		return nil, nil, err
	}
	if len(fromAccs) == 0 {
		return nil, nil, errors.New("gönderen müşterinin hesabı yok")
	}

	toAccs, err := repositorys.GetAccountsByCustomerID(toCustomerID)
	if err != nil {
		return nil, nil, err
	}
	if len(toAccs) == 0 {
		return nil, nil, errors.New("alıcı müşterinin hesabı yok")
	}

	// ✅ amount'u karşılayan bir gönderen hesabı seç
	var fromAccountID uint
	for _, a := range fromAccs {
		if a.Balance >= amount {
			fromAccountID = a.ID
			break
		}
	}
	if fromAccountID == 0 {
		return nil, nil, errors.New("yetersiz bakiye (müşterinin hiçbir hesabı bu tutarı karşılamıyor)")
	}

	// ✅ alıcı: ilk hesap
	toAccountID := toAccs[0].ID
	if toAccountID == 0 {
		return nil, nil, errors.New("alıcı hesap bulunamadı")
	}

	if fromAccountID == toAccountID {
		return nil, nil, errors.New("aynı hesaba transfer yapılamaz")
	}

	return repositorys.Transfer(fromAccountID, toAccountID, amount)
}
