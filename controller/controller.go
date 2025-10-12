package controller

import (
	"fianzy/models"

	"github.com/gin-gonic/gin"
)

func CreateCredit(c *gin.Context) {
	obj := models.Transaction{}
	if err := c.ShouldBindQuery(&obj); err != nil {
		c.JSON(400, gin.H{"message": "Invalid query"})
	}
	switch obj.Type {
	case "purchase":
		CreatePurchase(*obj.Purchase)
	case "investment":
		CreateInvestment(*obj.Investment)
	case "lend":
		CreateLend(*obj.Lend)
	case "borrow":
		CreateBorrow(*obj.Borrow)
	}
	c.JSON(200, gin.H{"message": "Credit created successfully"})
}

func CreateDebit(c *gin.Context) {
	obj := models.Transaction{}
	if err := c.ShouldBindQuery(&obj); err != nil {
		c.JSON(400, gin.H{"message": "Invalid query"})
	}
	switch obj.Type {
	case "purchase":
		CreatePurchase(*obj.Purchase)
	case "investment":
		CreateInvestment(*obj.Investment)
	case "lend":
		CreateLend(*obj.Lend)
	case "borrow":
		CreateBorrow(*obj.Borrow)
	}
	c.JSON(200, gin.H{"message": "Debit created successfully"})
}
