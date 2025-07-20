package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhpferreira/go-bitbank/services"
)

func CreateUser(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	wallet, err := services.CreateWallet(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wallet_id":   wallet.ID,
		"admin_key":   wallet.AdminKey,
		"invoice_key": wallet.InvoiceKey,
	})
}

func GenerateInvoice(c *gin.Context) {
	var req struct {
		Amount     int64  `json:"amount"`
		Memo       string `json:"memo"`
		InvoiceKey string `json:"invoice_key"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	invoice, err := services.GenerateInvoice(req.InvoiceKey, req.Amount, req.Memo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to generate invoince"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}
