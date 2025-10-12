package postgres

import (
	"fianzy/config"
	"fianzy/models"
)

func CreateInvestment(obj models.Investment) {
	config.DB.Create(&obj)
}
