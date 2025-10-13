package models

type Transaction struct {
	Type       string      `json:"type"` // purchase | borrow | lend | investment
	Purchase   *Purchase   `json:"purchase,omitempty"`
	Investment *Investment `json:"investment,omitempty"`
	Lend       *Lend       `json:"lend,omitempty"`
	Borrow     *Borrow     `json:"borrow,omitempty"`
	Bank       *Bank       `json:"bank"`
}
type Purchase struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	ItemName      string  `json:"item_name"`
	Amount        float64 `json:"amount"`
	PurchaseDate  int64   `json:"purchase_date"`
	PaymentMethod string  `json:"payment_method"`

	IsLinkedPurchase bool   `json:"is_linked_purchase"`
	LinkedFundID     uint   `json:"linked_fund_id"`
	LinkedFundType   string `json:"linked_fund_type"` //investment, lend
	RepaymentDate    int64  `json:"repayment_date"`
	IsRepaid         bool   `json:"is_repaid"`
	CreatedAt        int64
}

type Borrow struct {
	ID             uint    `json:"id" gorm:"primaryKey"`
	Type           string  `json:"type"` //card/loan/person
	PersonName     string  `json:"person_name"`
	Amount         float64 `json:"amount"`
	BorrowedDate   int64   `json:"borrowed_date"`
	RepaymentDate  int64   `json:"repayment_date"`
	IsRepaid       bool    `json:"is_repaid"`
	IsLinkedFund   bool    `json:"is_linked_fund"`
	LinkedFundID   uint    `json:"linked_fund_id"`
	LinkedFundType string  `json:"linked_fund_type"` //investment, lend
	Notes          string  `json:"notes"`
	CreatedAt      int64
}

type Investment struct {
	ID               uint    `json:"id" gorm:"primaryKey"`
	IsLinkedFund     bool    `json:"is_linked_fund"`
	LinkedFundType   string  `json:"linked_fund_type"`
	LinkedPurchaseID uint    `json:"linked_purchase_id"`
	LinkedBorrowID   uint    `json:"linked_borrow_id"`
	FundName         string  `json:"fund_name"`
	Amount           float64 `json:"amount"`
	ExpectedReturn   float64 `json:"expected_return"`
	InvestedDate     int64   `json:"invested_date"`
	CreatedAt        int64
}

type Lend struct {
	ID                 uint    `json:"id" gorm:"primaryKey"`
	IsLinkedFund       bool    `json:"is_linked_fund"`
	LinkedPurchaseID   uint    `json:"linked_purchase_id"`
	PersonName         string  `json:"person_name"`
	Amount             float64 `json:"amount"`
	GivenDate          int64   `json:"given_date"`
	ExpectedReturnDate int64   `json:"expected_return_date"`
	IsReturned         bool    `json:"is_returned"`
	Notes              string  `json:"notes"`
	CreatedAt          int64
}

type Bank struct {
	ID              uint    `json:"id" gorm:"primaryKey"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"` //credit or debit
	AccountType     string  `json:"account_type"`     //bank_account or credit_card
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
