package services

import "github.com/beyza/go-bank-simulator/repositorys"

type SeedUser struct {
	Name   string
	Email  string
	Amount float64
}

func SeedUsers(users []SeedUser) error {
	for _, u := range users {
		if err := repositorys.CreateUserWithAccount(u.Name, u.Email, u.Amount); err != nil {
			return err
		}
	}
	return nil
}
