package models

import "time"

type Transaction struct {
	Type       string      `json:"type"` // purchase | borrow | lend | investment
	Purchase   *Purchase   `json:"purchase,omitempty"`
	Investment *Investment `json:"investment,omitempty"`
	Lend       *Lend       `json:"lend,omitempty"`
	Borrow     *Borrow     `json:"borrow,omitempty"`
}
type Purchase struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	ItemName      string    `json:"item_name"`
	Amount        float64   `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PaymentMethod string    `json:"payment_method"`

	LinkedFundID   uint      `json:"linked_fund_id"`
	LinkedFundType string    `json:"linked_fund_type"` //investment, lend
	RepaymentDate  time.Time `json:"repayment_date"`
	IsRepaid       bool      `json:"is_repaid"`
	CreatedAt      time.Time
}

type Borrow struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	PersonName     string    `json:"person_name"`
	Amount         float64   `json:"amount"`
	BorrowedDate   time.Time `json:"borrowed_date"`
	RepaymentDate  time.Time `json:"repayment_date"`
	IsRepaid       bool      `json:"is_repaid"`
	LinkedFundID   uint      `json:"linked_fund_id"`
	LinkedFundType string    `json:"linked_fund_type"` //investment, lend
	Notes          string    `json:"notes"`
	CreatedAt      time.Time
}

type Investment struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	IsLinkedFund     bool      `json:"is_linked_fund"`
	LinkedPurchaseID uint      `json:"linked_purchase_id"`
	FundName         string    `json:"fund_name"`
	Amount           float64   `json:"amount"`
	ExpectedReturn   float64   `json:"expected_return"`
	InvestedDate     time.Time `json:"invested_date"`
	CreatedAt        time.Time
}

type Lend struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	IsLinkedFund       bool      `json:"is_linked_fund"`
	LinkedPurchaseID   uint      `json:"linked_purchase_id"`
	PersonName         string    `json:"person_name"`
	Amount             float64   `json:"amount"`
	GivenDate          time.Time `json:"given_date"`
	ExpectedReturnDate time.Time `json:"expected_return_date"`
	IsReturned         bool      `json:"is_returned"`
	Notes              string    `json:"notes"`
	CreatedAt          time.Time
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
}
