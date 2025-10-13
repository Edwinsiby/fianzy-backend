package controller

import (
	"fianzy/config"
	"fianzy/models"
	"fianzy/postgres"
)

func CreatePurchase(obj models.Purchase) error {
	return postgres.CreatePurchase(obj)
}

func GetPurchases() error {
	var purchases []models.Purchase
	config.DB.Find(&purchases)
	return nil
}
