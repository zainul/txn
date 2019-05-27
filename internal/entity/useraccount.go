package entity

// Generate by Thor

import (
	"time"

	xsvalidator "github.com/zainul/txn/internal/pkg/validator"
)

//UserAccount ....
type UserAccount struct {
	AccountNumber int64     `json:"account_number" gorm:"column:account_number;"`
	Balance       float64   `json:"balance" gorm:"column:balance;"`
	Modified      time.Time `json:"modified" gorm:"column:modified;"`
}

// Validate Useraccount entity...
func (e *UserAccount) Validate() error {
	return xsvalidator.Validate(e)
}

func (e *UserAccount) TableName() string {
	return "user_account"
}
