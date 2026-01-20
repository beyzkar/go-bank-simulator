package handlers

import (
	"net/http"
	"strconv"

	"github.com/beyza/go-bank-simulator/services"
	"github.com/gin-gonic/gin"
)

type createCustomerReq struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
}

func CreateCustomer(c *gin.Context) {
	var req createCustomerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz json body"})
		return
	}

	// ✅ müşteri + başlangıç bakiyeli account oluştur
	customer, account, err := services.CreateCustomerWithAccount(req.Name, req.Email, req.Balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ ikisini birlikte döndür
	c.JSON(http.StatusCreated, gin.H{
		"customer": customer,
		"account":  account,
	})
}

func GetAllCustomers(c *gin.Context) {
	customers, err := services.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func GetCustomerByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz id"})
		return
	}

	customer, err := services.GetCustomerByID(uint(id64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz id"})
		return
	}

	if err := services.DeleteCustomer(uint(id64)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func SearchCustomers(c *gin.Context) {
	q := c.Query("q")
	customers, err := services.SearchCustomersByName(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}
