package models

type Transaction struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Type   string  `json:"type"` // "purchase", "borrow", "lend", "investment"
	Amount float64 `json:"amount"`
	Notes  string  `json:"notes"`

	// Type-specific fields (nullable / optional)
	Name            *string `json:"name"`
	PaymentMethod   *string `json:"payment_method"`   // bank: bank_account or credit_card
	TransactionType *string `json:"transaction_type"` // bank: credit/debit

	// Linking between transactions
	LinkedID   *uint   `json:"linked_id"`
	LinkedName *string `json:"linked_name"`

	IsSettled     *bool  `json:"is_settled"`
	RepaymentDate *int64 `json:"repayment_date"`

	CreatedAt int64 `json:"created_at"`
}

type Stats struct {
	Asset           float32 `json:"asset"`
	DebtLinkedAsset float32 `json:"debt_linked_asset"`
	Debt            float32 `json:"debt"`
	AssetLinkedDebt float32 `json:"asset_linked_debt"`
	Investment      float32 `json:"investment"`
	Purchase        float32 `json:"purchase"`
	Lent            float32 `json:"lent"`
	Borrow          float32 `json:"borrow"`
	BankBalance     float32 `json:"bank_balance"`
	CreditCardUsed  float32 `json:"credit_card_used"`
}
