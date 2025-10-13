package controller

import (
	"fianzy/config"
	"fianzy/models"
	"fianzy/postgres"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBorrow(obj models.Borrow) error {
	postgres.CreateBorrow(obj)
	return nil
}

func GetBorrow() error {
	var obj []models.Borrow
	config.DB.Find(&obj)
	return nil
}

func RepayBorrow(c *gin.Context) {
	borrowIDStr := c.Param("borrow_id")
	borrowID, err := strconv.Atoi(borrowIDStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid borrow_id"})
		return
	}

	if err := postgres.RepayBorrow(borrowID); err != nil {
		c.JSON(500, gin.H{"message": "failed to repay borrow", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Borrow marked as repaid"})
}
