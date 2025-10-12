package controller

import (
	"fianzy/config"
	"fianzy/models"
	"fianzy/postgres"

	"github.com/gin-gonic/gin"
)

func CreateInvestment(obj models.Investment) error {
	postgres.CreateInvestment(obj)
	return nil
}

func GetInvestments(c *gin.Context) error {
	var investments []models.Investment
	config.DB.Find(&investments)
	return nil
}
