package postgres

import (
	"fianzy/config"
	"fianzy/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateTransactions(obj models.Transaction) error {
	if err := config.DB.Create(&obj).Error; err != nil {
		return err
	}
	return nil
}

func GetTransactions(txType string, page, limit int) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	db := config.DB.Model(&models.Transaction{})
	db = db.Where("type = ?", txType)
	offset := (page - 1) * limit
	if err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&transactions).Error; err != nil {

		return transactions, 0, err
	}
	var total int64
	if err := config.DB.Model(&models.Transaction{}).
		Where("type = ?", txType).
		Count(&total).Error; err != nil {
		total = int64(len(transactions))
	}

	return transactions, total, nil
}

func GetStats() (models.Stats, error) {
	var stats models.Stats
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Transaction{}).
			Where("type = ? AND is_settled = ?", "investment", false).Select("COALESCE(SUM(amount),0)").
			Scan(&stats.Investment).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Transaction{}).
			Where("type = ? AND is_settled = ? AND linked_id > 0", "investment", false).Select("COALESCE(SUM(amount),0)").
			Scan(&stats.DebtLinkedAsset).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Transaction{}).
			Where("type = ? AND is_settled = ?", "borrow", false).Select("COALESCE(SUM(amount),0)").
			Scan(&stats.Debt).Error; err != nil {
			return err
		}
		subQuery := tx.Model(&models.Transaction{}).
			Select("linked_id").
			Where("type = ? AND is_settled = ? AND linked_id > 0", "investment", false)
		if err := tx.Model(&models.Transaction{}).
			Where("id IN (?) AND type = ?", subQuery, "borrow").
			Select("COALESCE(SUM(amount),0)").
			Scan(&stats.AssetLinkedDebt).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Transaction{}).
			Where("type = ?", "purchase").Select("COALESCE(SUM(amount),0)").
			Scan(&stats.Purchase).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Transaction{}).
			Where("type = ? AND is_settled = ?", "lend", false).Select("COALESCE(SUM(amount),0)").
			Scan(&stats.Lent).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Transaction{}).
			Where("type = ? AND is_settled = ?", "borrow", false).Select("COALESCE(SUM(amount),0)").
			Scan(&stats.Borrow).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Transaction{}).
			Where("payment_method = ?", "bank").
			Select(`COALESCE(SUM(CASE 
                WHEN transaction_type = 'credit' THEN amount
                WHEN transaction_type = 'debit' THEN -amount
                ELSE 0
                END
                ), 0)
            `).
			Scan(&stats.BankBalance).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Transaction{}).
			Where("payment_method = ?", "credit_card").
			Select(`COALESCE(SUM(CASE 
                WHEN transaction_type = 'debit' THEN amount
                WHEN transaction_type = 'credit' THEN -amount
                ELSE 0
                END
                ), 0)
            `).
			Scan(&stats.CreditCardUsed).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return stats, err
	}

	stats.Asset = stats.Investment + stats.Lent + stats.BankBalance
	return stats, nil
}

func LinkFunds(sourceID, targetID int) error {
	db := config.DB
	var borrowTx, investTx models.Transaction

	// Fetch borrow transaction
	if err := db.First(&borrowTx, sourceID).Error; err != nil {
		return fmt.Errorf("borrow transaction not found: %w", err)
	}

	// Fetch investment transaction
	if err := db.First(&investTx, targetID).Error; err != nil {
		return fmt.Errorf("investment transaction not found: %w", err)
	}

	// Link the investment to borrow
	investTx.LinkedID = &borrowTx.ID
	investTx.LinkedName = borrowTx.Name

	// Save the updated investment
	if err := db.Save(&investTx).Error; err != nil {
		return fmt.Errorf("failed to link investment to borrow: %w", err)
	}

	return nil
}

func Setlement(id int64) error {
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var t models.Transaction
	if err := tx.First(&t, id).Error; err != nil {
		return err
	}

	if (t.Type != "borrow" && t.Type != "lend") || *t.IsSettled {
		return fmt.Errorf("invalid type")
	}
	if err := tx.Model(&t).Update("is_settled", true).Error; err != nil {
		return err
	}
	transactionType := "credit"
	paymentMethod := "bank"
	setlement := true
	if t.Type == "borrow" {
		transactionType = "debit"
	}

	bankTx := models.Transaction{
		Type:            "bank",
		Name:            t.Name,
		Amount:          t.Amount,
		PaymentMethod:   &paymentMethod,
		TransactionType: &transactionType,
		CreatedAt:       time.Now().Unix(),
		IsSettled:       &setlement,
	}
	if err := tx.Create(&bankTx).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}
