package services

/*
	repository ile handler arasında köprü görevi görür
	business logic burada yazılır (bu sistem nasıl çalışmalı, hangi durumda neye izin verilir)
	database'e gitmek business logic katmanının görevi değildir
	kontrol yapan katman
	service = beyin
*/

import (
	"errors"
	"strings"

	"github.com/beyza/go-bank-simulator/models"
	"github.com/beyza/go-bank-simulator/repositorys"
)

// Yeni müşteri oluşturur (iş kuralları burada)
func CreateCustomer(name, email string) (*models.Customer, error) {
	if strings.TrimSpace(name) == "" {
		//strings.TrimSpace: baştaki ve sondaki boşlukları siler
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

	//service -> repositorys -> database

	if err != nil {
		return nil, err
	}

	return customer, nil
}
func CreateCustomerWithAccount(name, email string, balance float64) (*models.Customer, *models.Account, error) {

	// 1️⃣ Müşteriyi oluştur
	customer, err := CreateCustomer(name, email)
	if err != nil {
		return nil, nil, err
	}

	// 2️⃣ Hesap oluştur
	account := &models.Account{
		CustomerID: customer.ID,
		Balance:    balance,
	}

	if err := repositorys.CreateAccount(account); err != nil {
		return nil, nil, err
	}

	return customer, account, nil
}

// ID ile müşteri getirir
func GetCustomerByID(id uint) (*models.Customer, error) {
	if id == 0 {
		return nil, errors.New("geçersiz müşteri id")
	}

	return repositorys.GetCustomerByID(id)
}

func GetCustomerByName(name string) (*models.Customer, error) {
	customer, err := repositorys.FindCustomerByName(name)
	if err != nil {
		return nil, errors.New("müşteri bulunamadı")
	}
	return customer, nil
}

// Tüm müşterileri getirir
func GetAllCustomers() ([]models.Customer, error) {
	return repositorys.GetAllCustomers()
}

func DeleteCustomer(id uint) error {
	// 1. Hesapları sil
	if err := repositorys.DeleteAccountsByCustomerID(id); err != nil {
		return err
	}

	// 2. Müşteriyi sil
	return repositorys.DeleteCustomer(id)
}
func SearchCustomersByName(q string) ([]models.Customer, error) {
	if strings.TrimSpace(q) == "" {
		return []models.Customer{}, nil
	}
	return repositorys.SearchCustomersByName(q)
}

func GetCustomerByExactName(name string) (*models.Customer, error) {
	list, err := repositorys.FindCustomerByExactName(name)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("müşteri bulunamadı")
	}
	if len(list) > 1 {
		return nil, errors.New("aynı isimde birden fazla müşteri var (ID ile seç)")
	}
	return &list[0], nil
}
