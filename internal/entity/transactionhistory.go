package entity

// Generate by Thor

import (
	"time"

	xsvalidator "github.com/zainul/txn/internal/pkg/validator"
)

//TransactionHistory ....
type TransactionHistory struct {
	TxID              string    `json:"tx_id" gorm:"column:tx_id;"`
	AccountNumber     int64     `json:"account_number" gorm:"column:account_number;"`
	TransactionStatus int       `json:"transaction_status" gorm:"column:transaction_status;"`
	DrCr              string    `json:"dr_cr" gorm:"column:dr_cr;"`
	PreviousBalance   float64   `json:"previous_balance" gorm:"column:previous_balance;"`
	EndBalance        float64   `json:"end_balance" gorm:"column:end_balance;"`
	TransactionCode   string    `json:"transaction_code" gorm:"column:transaction_code;"`
	Information       string    `json:"information" gorm:"column:information;"`
	UserComment       string    `json:"user_comment" gorm:"column:user_comment;"`
	TransactionDate   time.Time `json:"transaction_date" gorm:"column:transaction_date;"`
	Amount            float64   `json:"amount" gorm:"column:amount;"`
	OpponentAccount   int64     `json:"opponent_account" gorm:"column:opponent_account;"`
}

// Validate Transactionhistory entity...
func (e *TransactionHistory) Validate() error {
	return xsvalidator.Validate(e)
}

func (e *TransactionHistory) TableName() string {
	return "transaction_history"
}
