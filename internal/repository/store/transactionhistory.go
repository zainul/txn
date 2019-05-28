package store

import (
	"github.com/jinzhu/gorm"
	"github.com/zainul/txn/internal/entity"
)

type TransactionHistoryStore struct {
	DB *gorm.DB
}

// NewTransactionHistoryStore ...
func NewTransactionHistoryStore(conn *gorm.DB) *TransactionHistoryStore {
	return &TransactionHistoryStore{
		DB: conn,
	}
}

func (s *TransactionHistoryStore) Insert(e entity.TransactionHistory) error {
	return s.DB.Create(&e).Error
}
