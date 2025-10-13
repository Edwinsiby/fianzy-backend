package controller

import (
	"fianzy/config"
	"fianzy/models"
	"fianzy/postgres"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {
	data := postgres.GetStats()
	c.JSON(200, data)
}

func GetTransactions(c *gin.Context) {
	txType := c.Query("type") // e.g. "purchase", "borrow", "lend", "investment"

	switch txType {
	case "purchase":
		var purchases []models.Purchase
		if err := config.DB.Find(&purchases).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, purchases)

	case "investment":
		var investments []models.Investment
		if err := config.DB.Find(&investments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, investments)

	case "lend":
		var lends []models.Lend
		if err := config.DB.Find(&lends).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, lends)

	case "borrow":
		var borrows []models.Borrow
		if err := config.DB.Find(&borrows).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, borrows)

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing type parameter"})
	}
}

func LinkFunds(c *gin.Context) {
	linkType := c.Query("type")
	sourceIDStr := c.Query("source_id")
	targetIDStr := c.Query("target_id")

	sourceID, err := strconv.Atoi(sourceIDStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid source_id"})
		return
	}

	targetID, err := strconv.Atoi(targetIDStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid target_id"})
		return
	}

	if err := postgres.LinkFunds(linkType, sourceID, targetID); err != nil {
		c.JSON(500, gin.H{"message": "failed to link", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Link successful"})
}

func CreateBank(obj models.Bank) error {
	return postgres.CreateBank(obj)
}
