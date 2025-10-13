package postgres

import (
	"fianzy/config"
	"fianzy/models"
)

func CreateInvestment(obj models.Investment) error {
	if err := config.DB.Create(&obj).Error; err != nil {
		return err
	}
	return nil
}
