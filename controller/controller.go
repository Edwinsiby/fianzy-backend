package controller

import (
	"fianzy/models"
	"fianzy/postgres"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTransactions(c *gin.Context) {
	obj := models.Transaction{}
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(400, gin.H{"message": "Invalid JSON body", "error": err.Error()})
		return
	}
	if err := postgres.CreateTransactions(obj); err != nil {
		c.JSON(400, gin.H{"message": "", "error": err.Error()})
	}

	c.JSON(200, gin.H{"message": "Credit created successfully"})
	return
}

func GetTransactions(c *gin.Context) {
	txType := c.Query("type") // e.g. "purchase", "borrow", "lend", "investment"
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")
	if txType == "" {
		c.JSON(400, "invalid type")
	}
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}

	transactions, total, err := postgres.GetTransactions(txType, page, limit)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  transactions,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

func GetStats(c *gin.Context) {
	data, err := postgres.GetStats()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, data)
}

func LinkFunds(c *gin.Context) {
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

	if err := postgres.LinkFunds(sourceID, targetID); err != nil {
		c.JSON(500, gin.H{"message": "failed to link", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Link successful"})
}

func Setlement(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(400, gin.H{"error": "Missing transaction ID"})
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid transaction ID"})
		return
	}

	if err := postgres.Setlement(id); err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, nil)
}
