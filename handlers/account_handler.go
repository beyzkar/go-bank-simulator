package handlers

import (
	"net/http"
	"strconv"

	"github.com/beyza/go-bank-simulator/services"
	"github.com/gin-gonic/gin"
)

//create account

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

//get accounts by customer id

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

//get account by customer name  URL: /accounts/by-customer-name/:name

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
	lastTx, _ := services.GetLastTransactionByAccountID(account.ID)

	c.JSON(http.StatusOK, gin.H{
		"account_id":  account.ID,
		"customer":    customer.Name,
		"balance":     account.Balance,
		"last_action": lastTx,
	})
}

//get account by id

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

// delete account
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
