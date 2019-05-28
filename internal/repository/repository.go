package repository

import "github.com/zainul/txn/internal/entity"

// UserAccount ...
type UserAccount interface {
	Insert(p entity.UserAccount) error
	UserAccountBy(field string, value interface{}) ([]entity.UserAccount, error)
}

// TransactionHistory ...
type TransactionHistory interface {
}

// TransactionLog ...
type TransactionLog interface {
	Transfer(e []entity.TransactionLog) error
}
