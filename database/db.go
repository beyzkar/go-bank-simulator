package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/beyza/go-bank-simulator/models"
)

var DB *gorm.DB //uygulamada aktif bir veritabanı bağlantısı olacak ve ismi de DB olacak

func Init() { //veritabanını başlatan fonksiyon
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{}) //app.db dosyasını aç, SQLite veritabanına bağlan
	/*
		sqlite.Open("app.db") → dosyayı aç
		gorm.Open(...) → ORM ile bağlan
		db → artık veritabanı bağlantın
	*/
	if err != nil {
		log.Fatal("DB acilamadi: ", err) //log basar, programı direkt kapatır
	}

	// ✅ Modellerden tabloları üret
	if err := db.AutoMigrate(
		&models.Customer{},
		&models.Account{},
		&models.Transaction{},
	); err != nil {
		log.Fatal("AutoMigrate hatasi: ", err)
	}

	DB = db
	/*
		localde db değişkeni var, global DB değişkenine atıyoruz
		Handler’da database.DB  -> handler: dış dünyadan gelen isterkler
		Service’de database.DB
	*/
}
