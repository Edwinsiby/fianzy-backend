package controller

import (
	"fianzy/models"
	"fianzy/postgres"
)

func CreateLend(obj models.Lend) error {
	postgres.CreateLend(obj)
	return nil
}
