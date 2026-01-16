package main

import (
	"log"

	"github.com/beyza/go-bank-simulator/database"
)

func main() {
	database.Init() // veritabanını hazırla
	/*
		database.Init()
		   ↓
		DB’ye bağlan
		   ↓
		customers tablosunu oluştur
		   ↓
		accounts tablosunu oluştur
		   ↓
		transactions tablosunu oluştur
		   ↓
		DB’yi hazır hale getir

	*/
	log.Println("DB hazir. Tablolar olusturuldu.")
	/*
		burada fmt yerine log kullanmamızın sebebi, log paketinin zaman damgası eklemesi ve daha profesyonel bir çıktı sağlamasıdır.
		bu sayede herhnagi bir hata olduğunda veya bilgi vermek istediğimizde, zaman bilgisiyle birlikte daha anlamlı loglar elde ederiz.
		log.println: kayıt tutar
	*/

}
