package postgres

import (
	"errors"
	"fianzy/config"
	"fianzy/models"
)

func CreateBorrow(obj models.Borrow) error {
	if err := config.DB.Create(&obj).Error; err != nil {
		return err
	}
	return nil
}

func RepayBorrow(borrowID int) error {
	db := config.DB
	tx := db.Begin()

	var borrow models.Borrow
	if err := tx.First(&borrow, borrowID).Error; err != nil {
		tx.Rollback()
		return errors.New("borrow not found")
	}

	// ✅ Mark the borrow as repaid
	borrow.IsRepaid = true
	if err := tx.Save(&borrow).Error; err != nil {
		tx.Rollback()
		return err
	}

	// ✅ Optionally mark linked fund as "settled"
	if borrow.LinkedFundID != 0 {
		switch borrow.LinkedFundType {
		case "investment":
			var inv models.Investment
			if err := tx.First(&inv, borrow.LinkedFundID).Error; err == nil {
				// example: set a flag like IsUsedForRepayment
				inv.IsLinkedFund = false
				tx.Save(&inv)
			}
		case "purchase":
			var p models.Purchase
			if err := tx.First(&p, borrow.LinkedFundID).Error; err == nil {
				p.IsLinkedPurchase = false
				tx.Save(&p)
			}
		}
	}

	return tx.Commit().Error
}
