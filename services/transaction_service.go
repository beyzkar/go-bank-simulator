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
		return nil, errors.New("tutar 0'dan büyük olmalı")
	}
	return repositorys.Deposit(accountID, amount)
}

func Withdraw(accountID uint, amount float64) (*models.Transaction, error) {
	if accountID == 0 {
		return nil, errors.New("geçersiz account id")
	}
	if amount <= 0 {
		return nil, errors.New("tutar 0'dan büyük olmalı")
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

	// DB transaction + bakiye kontrolü repository içinde
	return repositorys.Transfer(fromAccountID, toAccountID, amount)
}

// =========================
// TRANSFER (CustomerID ile)
// ✅ En doğru yaklaşım: seçim ve bakiye kontrolü DB transaction içinde
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

	// ✅ Burada hesapları çekip seçmeye çalışma.
	// Çünkü “doğru hesap seçimi + yeterli bakiye kontrolü” DB transaction içinde yapılmalı.
	return repositorys.TransferByCustomerID(fromCustomerID, toCustomerID, amount)
}
