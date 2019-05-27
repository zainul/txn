package repository

import "github.com/zainul/txn/internal/entity"

// UserAccount ...
type UserAccount interface {
	Insert(p entity.UserAccount) error
	Update(p entity.UserAccount) error
	UpdateBy(e entity.UserAccount, param map[string]interface{}) error
	Delete(id interface{}) error
	UserAccountBy(field string, value interface{}) ([]entity.UserAccount, error)
	RawQuery(sql string, result interface{}) error
	RawExec(query string, param ...interface{}) error
}

// TransactionHistory ...
type TransactionHistory interface {
	Insert(p entity.TransactionHistory) error
	Update(p entity.TransactionHistory) error
	UpdateBy(e entity.TransactionHistory, param map[string]interface{}) error
	Delete(id interface{}) error
	TransactionHistoryBy(field string, value interface{}) ([]entity.TransactionHistory, error)
	RawQuery(sql string, result interface{}) error
	RawExec(query string, param ...interface{}) error
}

// TransactionLog ...
type TransactionLog interface {
	Insert(p entity.TransactionLog) error
	Update(p entity.TransactionLog) error
	UpdateBy(e entity.TransactionLog, param map[string]interface{}) error
	Delete(id interface{}) error
	TransactionLogBy(field string, value interface{}) ([]entity.TransactionLog, error)
	RawQuery(sql string, result interface{}) error
	RawExec(query string, param ...interface{}) error
	Transfer(e []entity.TransactionLog) error
}
