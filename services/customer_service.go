package services

import (
	"errors"
	"strings"

	"github.com/beyza/go-bank-simulator/models"
	"github.com/beyza/go-bank-simulator/repositorys"
)

// Yeni müşteri oluşturur (iş kuralları burada)
func CreateCustomer(name, email string) (*models.Customer, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("isim boş olamaz")
	}

	if strings.TrimSpace(email) == "" {
		return nil, errors.New("email boş olamaz")
	}

	customer := &models.Customer{
		Name:  name,
		Email: email,
	}

	err := repositorys.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// ID ile müşteri getirir
func GetCustomerByID(id uint) (*models.Customer, error) {
	if id == 0 {
		return nil, errors.New("geçersiz müşteri id")
	}

	return repositorys.GetCustomerByID(id)
}

// Tüm müşterileri getirir
func GetAllCustomers() ([]models.Customer, error) {
	return repositorys.GetAllCustomers()
}

// Müşteri siler
func DeleteCustomer(id uint) error {
	if id == 0 {
		return errors.New("geçersiz müşteri id")
	}

	return repositorys.DeleteCustomer(id)
}
