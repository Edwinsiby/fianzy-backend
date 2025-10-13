package postgres

import (
	"errors"
	"log"

	"fianzy/config"
	"fianzy/models"
)

func GetStats() models.Stats {
	var stats models.Stats

	// --- Investment Total ---
	if err := config.DB.Model(&models.Investment{}).
		Select("COALESCE(SUM(amount),0)").
		Scan(&stats.Investment).Error; err != nil {
		log.Println("Error fetching investment total:", err)
	}

	// --- Purchase Total ---
	if err := config.DB.Model(&models.Purchase{}).
		Select("COALESCE(SUM(amount),0)").
		Scan(&stats.Purchase).Error; err != nil {
		log.Println("Error fetching purchase total:", err)
	}

	// --- Lent Total ---
	if err := config.DB.Model(&models.Lend{}).
		Select("COALESCE(SUM(amount),0)").
		Scan(&stats.Lent).Error; err != nil {
		log.Println("Error fetching lent total:", err)
	}

	// --- Borrow Total ---
	if err := config.DB.Model(&models.Borrow{}).
		Select("COALESCE(SUM(amount),0)").
		Scan(&stats.Borrow).Error; err != nil {
		log.Println("Error fetching borrow total:", err)
	}

	// --- Bank Balance ---
	if err := config.DB.Model(&models.Bank{}).
		Select(`
		COALESCE(SUM(
			CASE 
				WHEN transaction_type = 'credit' THEN amount 
				WHEN transaction_type = 'debit' THEN -amount 
				ELSE 0 
			END
		), 0) as bank_balance`).
		Where("account_type = ?", "bank_account").
		Scan(&stats.BankBalance).Error; err != nil {
		log.Println("Error fetching bank balance:", err)
	}

	if err := config.DB.Model(&models.Bank{}).
		Select(`
		COALESCE(SUM(
			CASE 
				WHEN transaction_type = 'debit' THEN amount 
				ELSE 0 
			END
		), 0) as credit_card_used`).
		Where("account_type = ?", "credit_card").
		Scan(&stats.CreditCardUsed).Error; err != nil {
		log.Println("Error fetching credit used:", err)
	}

	// --- DebtLinkedAsset = purchases made with credit or linked borrow ---
	if err := config.DB.Model(&models.Purchase{}).
		Where("is_linked_purchase = ?", true).
		Select("COALESCE(SUM(amount),0)").
		Scan(&stats.DebtLinkedAsset).Error; err != nil {
		log.Println("Error fetching debt linked assets:", err)
	}

	// --- AssetLinkedDebt = borrowings linked to investment/lend ---
	if err := config.DB.Model(&models.Investment{}).
		Where("is_linked_fund = ?", true).
		Select("COALESCE(SUM(amount),0)").
		Scan(&stats.AssetLinkedDebt).Error; err != nil {
		log.Println("Error fetching asset linked debt:", err)
	}

	// --- Derived totals ---
	stats.Asset = stats.Investment + stats.Lent + stats.BankBalance
	stats.Debt = stats.Borrow

	return stats
}

func LinkFunds(linkType string, sourceID, targetID int) error {
	db := config.DB
	tx := db.Begin()

	switch linkType {

	// 🛒 Purchase funded by Borrow (e.g. credit card purchase)
	case "purchase_to_borrow":
		var purchase models.Purchase
		var borrow models.Borrow

		if err := tx.First(&purchase, sourceID).Error; err != nil {
			tx.Rollback()
			return errors.New("purchase not found")
		}
		if err := tx.First(&borrow, targetID).Error; err != nil {
			tx.Rollback()
			return errors.New("borrow not found")
		}

		purchase.IsLinkedPurchase = true
		purchase.LinkedFundID = borrow.ID
		purchase.LinkedFundType = "borrow"

		if err := tx.Save(&purchase).Error; err != nil {
			tx.Rollback()
			return err
		}

	// 💰 Investment to repay Borrow (e.g. invest to clear credit)
	case "investment_to_borrow":
		var investment models.Investment
		var borrow models.Borrow

		if err := tx.First(&investment, sourceID).Error; err != nil {
			tx.Rollback()
			return errors.New("investment not found")
		}
		if err := tx.First(&borrow, targetID).Error; err != nil {
			tx.Rollback()
			return errors.New("borrow not found")
		}

		investment.IsLinkedFund = true
		// Indicate this investment is for paying back the borrow
		investment.LinkedBorrowID = uint(borrow.ID) // optional: rename field if you want cleaner structure
		if err := tx.Save(&investment).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Link back for clarity (optional)
		borrow.LinkedFundID = investment.ID
		borrow.LinkedFundType = "investment"
		if err := tx.Save(&borrow).Error; err != nil {
			tx.Rollback()
			return err
		}

	// 📝 Borrow funds used for investment (e.g. borrowed to invest)
	case "borrow_to_investment":
		var borrow models.Borrow
		var investment models.Investment

		if err := tx.First(&borrow, sourceID).Error; err != nil {
			tx.Rollback()
			return errors.New("borrow not found")
		}
		if err := tx.First(&investment, targetID).Error; err != nil {
			tx.Rollback()
			return errors.New("investment not found")
		}

		borrow.LinkedFundID = investment.ID
		borrow.LinkedFundType = "investment"
		if err := tx.Save(&borrow).Error; err != nil {
			tx.Rollback()
			return err
		}

		investment.IsLinkedFund = true
		if err := tx.Save(&investment).Error; err != nil {
			tx.Rollback()
			return err
		}
	case "borrow_to_purchase":
		var borrow models.Borrow
		var purchase models.Purchase

		if err := tx.First(&borrow, sourceID).Error; err != nil {
			tx.Rollback()
			return errors.New("borrow not found")
		}
		if err := tx.First(&purchase, targetID).Error; err != nil {
			tx.Rollback()
			return errors.New("investment not found")
		}

		borrow.LinkedFundID = purchase.ID
		borrow.LinkedFundType = "purchase"
		if err := tx.Save(&borrow).Error; err != nil {
			tx.Rollback()
			return err
		}

		borrow.IsLinkedFund = true
		if err := tx.Save(&borrow).Error; err != nil {
			tx.Rollback()
			return err
		}

	default:
		tx.Rollback()
		return errors.New("invalid link type")
	}

	return tx.Commit().Error
}

func CreateBank(obj models.Bank) error {
	if err := config.DB.Create(&obj).Error; err != nil {
		return err
	}
	return nil
}
