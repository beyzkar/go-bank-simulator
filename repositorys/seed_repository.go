package repositorys

import (
	"time"

	"github.com/beyza/go-bank-simulator/database"
	"github.com/beyza/go-bank-simulator/models"
	"gorm.io/gorm"
)

func CreateUserWithAccount(name, email string, amount float64) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		customer := models.Customer{
			Name:  name,
			Email: email,
		}
		if err := tx.Create(&customer).Error; err != nil {
			return err
		}

		account := models.Account{
			CustomerID: customer.ID,
			Balance:    amount,
		}
		if err := tx.Create(&account).Error; err != nil {
			return err
		}

		t := models.Transaction{
			AccountID: account.ID,
			Type:      "deposit",
			Amount:    amount,
			CreatedAt: time.Now(),
		}
		if err := tx.Create(&t).Error; err != nil {
			return err
		}

		return nil
	})
}
