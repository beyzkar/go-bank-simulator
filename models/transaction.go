package models

import "time"

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	AccountID uint
	Type      string // deposit, withdraw, transfer
	Amount    float64
	CreatedAt time.Time
}
