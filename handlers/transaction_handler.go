package handlers

import (
	"net/http"
	"strconv"

	"github.com/beyza/go-bank-simulator/services"
	"github.com/gin-gonic/gin"
)

type moneyReq struct {
	Amount float64 `json:"amount"`
}

// POST /accounts/:id/deposit
func Deposit(c *gin.Context) {
	acc64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || acc64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz accountId"})
		return
	}

	var req moneyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz json body"})
		return
	}

	tx, err := services.Deposit(uint(acc64), req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tx)
}

// POST /accounts/:id/withdraw
func Withdraw(c *gin.Context) {
	acc64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || acc64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz accountId"})
		return
	}

	var req moneyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz json body"})
		return
	}

	tx, err := services.Withdraw(uint(acc64), req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tx)
}

// GET /transactions/:id
func GetTransactionByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz transaction id"})
		return
	}

	tx, err := services.GetTransactionByID(uint(id64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tx)
}

// GET /accounts/:id/transactions
func GetTransactionsByAccountID(c *gin.Context) {
	acc64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || acc64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz accountId"})
		return
	}

	txs, err := services.GetTransactionsByAccountID(uint(acc64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, txs)
}

type transferRequest struct {
	FromAccountID uint    `json:"fromAccountId"`
	ToAccountID   uint    `json:"toAccountId"`
	Amount        float64 `json:"amount"`
}

// POST /accounts/transfer
func Transfer(c *gin.Context) {
	var req transferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz json body"})
		return
	}

	txOut, txIn, err := services.Transfer(req.FromAccountID, req.ToAccountID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transferOut": txOut,
		"transferIn":  txIn,
	})
}

// ✅ CustomerID ile transfer
type transferByCustomerRequest struct {
	FromCustomerID uint    `json:"fromCustomerId"`
	ToCustomerID   uint    `json:"toCustomerId"`
	Amount         float64 `json:"amount"`
}

// POST /transfer/by-customer
func TransferByCustomerID(c *gin.Context) {
	var req transferByCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz json body"})
		return
	}

	txOut, txIn, err := services.TransferByCustomerID(req.FromCustomerID, req.ToCustomerID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transferOut": txOut,
		"transferIn":  txIn,
	})
}
