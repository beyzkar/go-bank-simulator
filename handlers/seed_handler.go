package handlers

import (
	"net/http"

	"github.com/beyza/go-bank-simulator/services"
	"github.com/gin-gonic/gin"
)

type SeedUserReq struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Amount float64 `json:"amount"`
}

func SeedDatabase(c *gin.Context) {
	var req []SeedUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geçersiz json"})
		return
	}

	// ✅ Handler tipini Service tipine çevir
	users := make([]services.SeedUser, 0, len(req))
	for _, u := range req {
		users = append(users, services.SeedUser{
			Name:   u.Name,
			Email:  u.Email,
			Amount: u.Amount,
		})
	}

	if err := services.SeedUsers(users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seed işlemi başarılı"})
}
