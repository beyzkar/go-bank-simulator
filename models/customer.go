package models

import "time"

type Customer struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time
}
