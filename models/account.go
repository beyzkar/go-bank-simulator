package models

import "time"

type Account struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint
	Balance    float64
	CreatedAt  time.Time
}
