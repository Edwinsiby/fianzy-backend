package controller

import (
	"fianzy/models"

	"github.com/gin-gonic/gin"
)

func CreateTransactions(c *gin.Context) {
	obj := models.Transaction{}
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(400, gin.H{"message": "Invalid JSON body", "error": err.Error()})
		return
	}
	switch obj.Type {
	case "purchase":
		if err := CreatePurchase(*obj.Purchase); err != nil {
			c.JSON(500, err)
			return
		}
	case "investment":
		if err := CreateInvestment(*obj.Investment); err != nil {
			c.JSON(500, err)
			return
		}
	case "lend":
		if err := CreateLend(*obj.Lend); err != nil {
			c.JSON(500, err)
			return
		}
	case "borrow":
		if err := CreateBorrow(*obj.Borrow); err != nil {
			c.JSON(500, err)
			return
		}
	case "bank":
		if err := CreateBank(*obj.Bank); err != nil {
			c.JSON(500, err)
			return
		}
	default:
		c.JSON(400, gin.H{"message": "Invalid query"})
		return
	}
	c.JSON(200, gin.H{"message": "Credit created successfully"})
	return
}
