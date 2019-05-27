package entity

// Generate by Thor

import (
	"time"

	xsvalidator "github.com/zainul/txn/internal/pkg/validator"
)

//TransactionLog ....
type TransactionLog struct {
	TxID                  string    `json:"tx_id" gorm:"column:tx_id;"`
	AccountNumber         int64     `json:"account_number" gorm:"column:account_number;"`
	Date                  time.Time `json:"date" gorm:"column:date;"`
	TransactionCode       string    `json:"transaction_code" gorm:"column:transaction_code;"`
	PreviousBalance       float64   `json:"previous_balance" gorm:"column:previous_balance;"`
	EndBalance            float64   `json:"end_balance" gorm:"column:end_balance;"`
	OpponentAccountNumber int64     `json:"opponent_account_number" gorm:"column:opponent_account_number;"`
	TransactionStatus     int       `json:"transaction_status" gorm:"column:transaction_status;"`
	DrCr                  string    `json:"dr_cr" gorm:"column:dr_cr;"`
	Information           string    `json:"information" gorm:"column:information;"`
	LastUpdate            time.Time `json:"last_update" gorm:"column:last_update;"`
	Amount                float64   `json:"amount" gorm:"column:amount;"`
}

// Validate Transactionlog entity...
func (e *TransactionLog) Validate() error {
	return xsvalidator.Validate(e)
}

func (e *TransactionLog) TableName() string {
	return "transaction_log"
}
