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

// Müşteri siler
func DeleteCustomer(id uint) error {
	if id == 0 { //bu ksıım business logic
		return errors.New("geçersiz müşteri id") // business logic değil, teknik işlem git veritabanından getir
	}

	return repositorys.DeleteCustomer(id)
}
