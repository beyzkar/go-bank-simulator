package main

import (
	"log"

	"github.com/beyza/go-bank-simulator/database"
	"github.com/beyza/go-bank-simulator/handlers"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.POST("/seed", handlers.SeedDatabase)

	// Customers
	r.POST("/customers", handlers.CreateCustomer)
	r.GET("/customers", handlers.GetAllCustomers)
	r.GET("/customers/:id", handlers.GetCustomerByID)
	r.DELETE("/customers/:id", handlers.DeleteCustomer)

	// Accounts
	r.POST("/accounts", handlers.CreateAccount)
	r.GET("/accounts/:id", handlers.GetAccountByID)
	r.DELETE("/accounts/:id", handlers.DeleteAccount)
	r.GET("/customers/:id/accounts", handlers.GetAccountsByCustomerID)

	//Müşteri ismiyle hesap sorgu
	r.GET("/accounts/by-customer-name/:name", handlers.GetAccountByCustomerName)

	// Transactions
	r.POST("/accounts/:id/deposit", handlers.Deposit)
	r.POST("/accounts/:id/withdraw", handlers.Withdraw)
	r.GET("/transactions/:id", handlers.GetTransactionByID)
	r.GET("/accounts/:id/transactions", handlers.GetTransactionsByAccountID)

	//Transfer
	r.POST("/accounts/transfer", handlers.Transfer)

	//log.Println("DB hazir. Tablolar olusturuldu.")
	/*
		burada fmt yerine log kullanmamızın sebebi, log paketinin zaman damgası eklemesi ve daha profesyonel bir çıktı sağlamasıdır.
		bu sayede herhnagi bir hata olduğunda veya bilgi vermek istediğimizde, zaman bilgisiyle birlikte daha anlamlı loglar elde ederiz.
		log.println: kayıt tutar
	*/

	log.Println("Server 8080 portunda çalışıyor...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
