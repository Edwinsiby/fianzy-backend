package postgres

import (
	"fianzy/config"
	"fianzy/models"
)

func CreateBorrow(obj models.Borrow) error {
	if err := config.DB.Create(&obj).Error; err != nil {
		return err
	}
	return nil
}
