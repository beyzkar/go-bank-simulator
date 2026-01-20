package repositorys

import (
	"github.com/beyza/go-bank-simulator/database"
	"github.com/beyza/go-bank-simulator/models"
)

// Yeni hesap oluşturur
func CreateAccount(account *models.Account) error {
	return database.DB.Create(account).Error
}

// ID ile hesap getirir
func GetAccountByID(id uint) (*models.Account, error) {
	var account models.Account
	err := database.DB.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// Bir müşterinin tüm hesaplarını getirir
func GetAccountsByCustomerID(customerID uint) ([]models.Account, error) {
	var accounts []models.Account
	err := database.DB.Where("customer_id = ?", customerID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
func FindAccountByCustomerID(customerID uint) (*models.Account, error) {
	var acc models.Account
	err := database.DB.Where("customer_id = ?", customerID).First(&acc).Error
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

// Hesabı günceller (balance vb. değişikliklerde kullanılır)
func UpdateAccount(account *models.Account) error {
	return database.DB.Save(account).Error
}

// Hesap siler
func DeleteAccount(id uint) error {
	return database.DB.Delete(&models.Account{}, id).Error
}
