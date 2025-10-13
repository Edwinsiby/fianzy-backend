package postgres

import (
	"fianzy/config"
	"fianzy/models"
)

func CreatePurchase(obj models.Purchase) error {
	if err := config.DB.Create(&obj).Error; err != nil {
		return err
	}
	return nil
}
