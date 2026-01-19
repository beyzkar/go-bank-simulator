package models

import "time"

type Customer struct {
	ID uint `gorm:"primaryKey"`
	//uint: auto increment
	//database tarafından otomatik olarak atanır
	Name      string
	Email     string
	CreatedAt time.Time
}
