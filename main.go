package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
}

func main() {
	// Bu dosya DB Browser ile açacağın dosya
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Tablo yoksa oluşturur (migration)
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	// Örnek insert
	u := User{Name: "Ali", Email: "ali@example.com"}
	if err := db.Create(&u).Error; err != nil {
		panic(err)
	}

	// Örnek select
	var users []User
	if err := db.Find(&users).Error; err != nil {
		panic(err)
	}

	fmt.Println("Users:", users)
}
