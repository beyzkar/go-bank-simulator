package repositorys

import (
	"errors"

	"github.com/beyza/go-bank-simulator/database"
	"github.com/beyza/go-bank-simulator/models"
	"gorm.io/gorm"
)

// ---- Okuma (Read) tarafı ----

func CreateTransaction(t *models.Transaction) error {
	return database.DB.Create(t).Error
}

func GetTransactionByID(id uint) (*models.Transaction, error) {
	var t models.Transaction
	if err := database.DB.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTransactionsByAccountID(accountID uint) ([]models.Transaction, error) {
	var txs []models.Transaction
	if err := database.DB.Where("account_id = ?", accountID).Order("id desc").Find(&txs).Error; err != nil {
		return nil, err
	}
	return txs, nil
}

// ---- Para hareketi (Deposit / Withdraw) ----
// Bu fonksiyonlar Account balance + Transaction kaydını TEK işlemde yapar.

func Deposit(accountID uint, amount float64) (*models.Transaction, error) {
	var createdTx *models.Transaction

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var account models.Account

		// Hesabı bul
		if err := tx.First(&account, accountID).Error; err != nil {
			return err
		}

		// Bakiye artır
		account.Balance += amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		// Transaction kaydı
		t := &models.Transaction{
			AccountID: accountID,
			Amount:    amount,
			Type:      "deposit",
		}

		if err := tx.Create(t).Error; err != nil {
			return err
		}

		createdTx = t
		return nil
	})

	if err != nil {
		return nil, err
	}
	return createdTx, nil
}

func Withdraw(accountID uint, amount float64) (*models.Transaction, error) {
	var createdTx *models.Transaction

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var account models.Account

		// Hesabı bul
		if err := tx.First(&account, accountID).Error; err != nil {
			return err
		}

		// Yeterli bakiye kontrolü (DB transaction içinde güvenli)
		if account.Balance < amount {
			return errors.New("yetersiz bakiye")
		}

		// Bakiye azalt
		account.Balance -= amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		// Transaction kaydı
		t := &models.Transaction{
			AccountID: accountID,
			Amount:    amount,
			Type:      "withdraw",
		}

		if err := tx.Create(t).Error; err != nil {
			return err
		}

		createdTx = t
		return nil
	})

	if err != nil {
		return nil, err
	}
	return createdTx, nil
}
