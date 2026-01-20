package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/beyza/go-bank-simulator/services"
	"github.com/gin-gonic/gin"
)

// Create account

type createAccountReq struct {
	CustomerID uint `json:"customerId"`
}

func CreateAccount(c *gin.Context) {
	var req createAccountReq

	if err := c.ShouldBindJSON(&req); err != nil || req.CustomerID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz json body"})
		return
	}

	account, err := services.CreateAccount(req.CustomerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

// Get accounts by customer id

func GetAccountsByCustomerID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz customer id"})
		return
	}

	accounts, err := services.GetAccountsByCustomerID(uint(id64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// Get account by customer name
// URL: /accounts/by-customer-name/:name

func GetAccountByCustomerName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "müşteri adı gerekli"})
		return
	}

	customer, err := services.GetCustomerByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "müşteri bulunamadı"})
		return
	}

	accounts, err := services.GetAccountsByCustomerID(customer.ID)
	if err != nil || len(accounts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "hesap bulunamadı"})
		return
	}

	account := accounts[0] // ilk hesabı al

	// Son işlem metni
	lastAction := "Henüz işlem yok"
	if tx, err := services.GetLastTransactionByAccountID(account.ID); err == nil && tx != nil {
		switch tx.Type {
		case "deposit":
			lastAction = fmt.Sprintf("%.0f TL yatırıldı", tx.Amount)
		case "withdraw":
			lastAction = fmt.Sprintf("%.0f TL çekildi", tx.Amount)
		case "transfer":
			lastAction = fmt.Sprintf("%.0f TL transfer edildi", tx.Amount)
		default:
			lastAction = fmt.Sprintf("Son işlem: %.0f TL (%s)", tx.Amount, tx.Type)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"customerName": customer.Name,
		"accountId":    account.ID,
		"customerId":   account.CustomerID,
		"balance":      account.Balance,
		"lastAction":   lastAction,
	})
}

// Get account details by account id
// URL: /accounts/:id/details

func GetAccountDetailsByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz account id"})
		return
	}

	accountID := uint(id64)

	// Hesabı getir
	account, err := services.GetAccountByID(accountID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "hesap bulunamadı"})
		return
	}

	// Müşteriyi getir
	customer, err := services.GetCustomerByID(account.CustomerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "müşteri bulunamadı"})
		return
	}

	// Son işlem metni
	lastAction := "Henüz işlem yok"
	if tx, err := services.GetLastTransactionByAccountID(account.ID); err == nil && tx != nil {
		switch tx.Type {
		case "deposit":
			lastAction = fmt.Sprintf("%.0f TL yatırıldı", tx.Amount)
		case "withdraw":
			lastAction = fmt.Sprintf("%.0f TL çekildi", tx.Amount)
		case "transfer":
			lastAction = fmt.Sprintf("%.0f TL transfer edildi", tx.Amount)
		default:
			lastAction = fmt.Sprintf("Son işlem: %.0f TL (%s)", tx.Amount, tx.Type)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"customerName": customer.Name,
		"accountId":    account.ID,
		"customerId":   account.CustomerID,
		"balance":      account.Balance,
		"lastAction":   lastAction,
	})
}

// Get account by id
// URL: /accounts/:id

func GetAccountByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz account id"})
		return
	}

	account, err := services.GetAccountByID(uint(id64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// Delete account

func DeleteAccount(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz account id"})
		return
	}

	if err := services.DeleteAccount(uint(id64)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
