package repositorys

import (
	"errors"
	"time"

	"github.com/beyza/go-bank-simulator/database"
	"github.com/beyza/go-bank-simulator/models"
	"gorm.io/gorm"
)

const eps = 0.000001

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

func FindLastTransactionByAccountID(accountID uint) (*models.Transaction, error) {
	var txm models.Transaction
	err := database.DB.
		Where("account_id = ?", accountID).
		Order("created_at desc").
		First(&txm).Error

	if err != nil {
		return nil, err
	}
	return &txm, nil
}

// ---- Para hareketi (Deposit / Withdraw) ----
// Bu fonksiyonlar Account balance + Transaction kaydını TEK işlemde yapar.

func Deposit(accountID uint, amount float64) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("tutar 0'dan büyük olmalı")
	}

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
	if amount <= 0 {
		return nil, errors.New("tutar 0'dan büyük olmalı")
	}

	var createdTx *models.Transaction

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var account models.Account

		// Hesabı bul
		if err := tx.First(&account, accountID).Error; err != nil {
			return err
		}

		// ✅ Float toleransı
		if account.Balance+eps < amount {
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

// ---- AccountID ile Transfer ----

func Transfer(fromID, toID uint, amount float64) (*models.Transaction, *models.Transaction, error) {
	if amount <= 0 {
		return nil, nil, errors.New("tutar 0'dan büyük olmalı")
	}

	db := database.DB
	if db == nil {
		return nil, nil, errors.New("db bağlantısı yok")
	}

	var txOut *models.Transaction
	var txIn *models.Transaction

	err := db.Transaction(func(tx *gorm.DB) error {
		var fromAcc models.Account
		if err := tx.First(&fromAcc, fromID).Error; err != nil {
			return errors.New("gönderen hesap bulunamadı")
		}

		var toAcc models.Account
		if err := tx.First(&toAcc, toID).Error; err != nil {
			return errors.New("alıcı hesap bulunamadı")
		}

		if fromAcc.ID == toAcc.ID {
			return errors.New("aynı hesaba transfer yapılamaz")
		}

		// ✅ Float toleransı
		if fromAcc.Balance+eps < amount {
			return errors.New("yetersiz bakiye")
		}

		// Bakiyeleri güncelle
		fromAcc.Balance -= amount
		toAcc.Balance += amount

		if err := tx.Save(&fromAcc).Error; err != nil {
			return err
		}
		if err := tx.Save(&toAcc).Error; err != nil {
			return err
		}

		now := time.Now()

		out := &models.Transaction{
			AccountID: fromAcc.ID,
			Type:      "transfer_out",
			Amount:    amount,
			CreatedAt: now,
		}
		in := &models.Transaction{
			AccountID: toAcc.ID,
			Type:      "transfer_in",
			Amount:    amount,
			CreatedAt: now,
		}

		if err := tx.Create(out).Error; err != nil {
			return err
		}
		if err := tx.Create(in).Error; err != nil {
			return err
		}

		txOut = out
		txIn = in
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return txOut, txIn, nil
}

// ---- CustomerID ile Transfer ----
// ✅ Gönderen tarafta: amount'u karşılayan hesabı seçer (balance desc içinde ilk eşleşen)
// ✅ Alıcı tarafta: ilk hesabı seçer

func TransferByCustomerID(fromCustomerID, toCustomerID uint, amount float64) (*models.Transaction, *models.Transaction, error) {
	if amount <= 0 {
		return nil, nil, errors.New("tutar 0'dan büyük olmalı")
	}

	db := database.DB
	if db == nil {
		return nil, nil, errors.New("db bağlantısı yok")
	}

	var txOut *models.Transaction
	var txIn *models.Transaction

	err := db.Transaction(func(tx *gorm.DB) error {
		// ✅ Gönderen: amount'u karşılayan hesabı seç
		var fromAcc models.Account
		if err := tx.
			Where("customer_id = ? AND balance >= ?", fromCustomerID, amount-eps).
			Order("balance desc").
			First(&fromAcc).Error; err != nil {
			// hiç bir hesap amount'u karşılamıyorsa
			return errors.New("yetersiz bakiye")
		}

		// ✅ Alıcı: ilk hesap
		var toAcc models.Account
		if err := tx.
			Where("customer_id = ?", toCustomerID).
			Order("id asc").
			First(&toAcc).Error; err != nil {
			return errors.New("alıcı müşterinin hesabı yok")
		}

		if fromAcc.ID == toAcc.ID {
			return errors.New("aynı hesaba transfer yapılamaz")
		}

		// ekstra güvenlik
		if fromAcc.Balance+eps < amount {
			return errors.New("yetersiz bakiye")
		}

		// Bakiyeleri güncelle
		fromAcc.Balance -= amount
		toAcc.Balance += amount

		if err := tx.Save(&fromAcc).Error; err != nil {
			return err
		}
		if err := tx.Save(&toAcc).Error; err != nil {
			return err
		}

		now := time.Now()

		out := &models.Transaction{
			AccountID: fromAcc.ID,
			Type:      "transfer_out",
			Amount:    amount,
			CreatedAt: now,
		}
		in := &models.Transaction{
			AccountID: toAcc.ID,
			Type:      "transfer_in",
			Amount:    amount,
			CreatedAt: now,
		}

		if err := tx.Create(out).Error; err != nil {
			return err
		}
		if err := tx.Create(in).Error; err != nil {
			return err
		}

		txOut = out
		txIn = in
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return txOut, txIn, nil
}
