package controller

import (
	"fianzy/config"
	"fianzy/models"
	"fianzy/postgres"
)

func CreateBorrow(obj models.Borrow) error {
	postgres.CreateBorrow(obj)
	return nil
}

func GetBorrow() error {
	var obj []models.Borrow
	config.DB.Find(&obj)
	return nil
}
