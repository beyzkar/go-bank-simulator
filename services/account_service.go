package services

import (
	"errors"

	"github.com/beyza/go-bank-simulator/models"
	"github.com/beyza/go-bank-simulator/repositorys"
)

// Yeni hesap açar (iş kuralları burada)
func CreateAccount(customerID uint) (*models.Account, error) {
	if customerID == 0 {
		return nil, errors.New("geçersiz customer id")
	}

	// ✅ İş kuralı: Hesap açmak için müşteri gerçekten var mı?
	_, err := repositorys.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}

	account := &models.Account{
		CustomerID: customerID,
		Balance:    0, // yeni hesap sıfır bakiye ile başlar
	}

	err = repositorys.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// ID ile hesap getirir
func GetAccountByID(id uint) (*models.Account, error) {
	if id == 0 {
		return nil, errors.New("geçersiz account id")
	}

	return repositorys.GetAccountByID(id)
}

// CustomerID ile hesapları getirir
func GetAccountsByCustomerID(customerID uint) ([]models.Account, error) {
	if customerID == 0 {
		return nil, errors.New("geçersiz customer id")
	}

	return repositorys.GetAccountsByCustomerID(customerID)
}

// Hesap siler
func DeleteAccount(id uint) error {
	if id == 0 {
		return errors.New("geçersiz account id")
	}

	return repositorys.DeleteAccount(id)
}
