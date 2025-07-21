package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizhpferreira/go-bitbank/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/create-user", handlers.CreateUser)
	r.POST("/create-invoice", handlers.CreateInvoice)

	r.Run(":8081")
}
