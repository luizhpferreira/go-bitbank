package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhpferreira/go-bitbank/services"
)

func CreateUser(c *gin.Context) {
	var userRequest struct {
		Username       string `json:"username"`
		Password       string `json:"password"`
		PasswordRepeat string `json:"password_repeat"`
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.CreateUser(userRequest.Username, userRequest.Password, userRequest.PasswordRepeat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func CreateInvoice(c *gin.Context) {
	var invoiceRequest struct {
		Out    bool   `json:"out"`
		Amount int    `json:"amount"`
		Unit   string `json:"unit"`
		Memo   string `json:"memo"`
	}

	if err := c.ShouldBindJSON(&invoiceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.CreateInvoice(invoiceRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice created successfully"})
}
