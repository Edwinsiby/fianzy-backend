package postgres

import (
	"fianzy/config"
	"fianzy/models"
)

func CreateLend(obj models.Lend) error {
	if err := config.DB.Create(&obj).Error; err != nil {
		return err
	}
	return nil
}
