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

	// Static dosyalar (JS, CSS)
	r.Static("/static", "./static")

	// HTML şablonları
	r.LoadHTMLGlob("templates/*")

	// Test endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Ana sayfa
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	// Seed (örnek veriler)
	r.POST("/seed", handlers.SeedDatabase)

	// Customers
	r.POST("/customers", handlers.CreateCustomer)
	r.GET("/customers", handlers.GetAllCustomers)
	r.GET("/customers/:id", handlers.GetCustomerByID)
	r.DELETE("/customers/:id", handlers.DeleteCustomer)
	r.GET("/customers/search", handlers.SearchCustomers)

	// Accounts
	r.POST("/accounts", handlers.CreateAccount)

	// Özel (isimle arama / detay) route'lar önce
	r.GET("/accounts/by-customer-name/:name", handlers.GetAccountByCustomerName)
	r.GET("/accounts/:id/details", handlers.GetAccountDetailsByID)

	// Genel ID route'ları sonra
	r.GET("/accounts/:id", handlers.GetAccountByID)
	r.DELETE("/accounts/:id", handlers.DeleteAccount)

	r.GET("/customers/:id/accounts", handlers.GetAccountsByCustomerID)

	// Transactions
	r.POST("/accounts/:id/deposit", handlers.Deposit)
	r.POST("/accounts/:id/withdraw", handlers.Withdraw)
	r.GET("/transactions/:id", handlers.GetTransactionByID)
	r.GET("/accounts/:id/transactions", handlers.GetTransactionsByAccountID)

	// Transfer
	r.POST("/accounts/transfer", handlers.Transfer)

	// ✅ CustomerID ile Transfer
	// NOT: Handler fonksiyon adın TransferByCustomer ise bu satır doğru.
	// Eğer TransferByCustomerID kullanacaksan handler adını da öyle yapmalısın.

	/*
		burada fmt yerine log kullanmamızın sebebi,
		log paketinin zaman damgası eklemesi ve
		daha profesyonel bir çıktı sağlamasıdır.
	*/

	log.Println("Server 8080 portunda çalışıyor...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
