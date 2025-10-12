package controller

import (
	"fianzy/config"
	"fianzy/models"
	"fianzy/postgres"
)

func CreatePurchase(obj models.Purchase) error {
	postgres.CreatePurchase(obj)
	return nil
}

func GetPurchases() error {
	var purchases []models.Purchase
	config.DB.Find(&purchases)
	return nil
}
