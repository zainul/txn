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

func (s *TransactionHistoryStore) Update(e entity.TransactionHistory) error {
	return s.DB.Table(e.TableName()).Where("tx_id = ?", e.TxID).Update(&e).Error
}

func (s *TransactionHistoryStore) UpdateBy(e entity.TransactionHistory, param map[string]interface{}) error {
	return s.DB.Table(e.TableName()).Where(param).Update(&e).Error
}

func (s *TransactionHistoryStore) Delete(id interface{}) error {
	return nil
}

func (s *TransactionHistoryStore) RawQuery(sql string, result interface{}) error {
	return s.DB.Raw(sql).Find(result).Error
}

func (s *TransactionHistoryStore) RawExec(query string, param ...interface{}) error {
	return s.DB.Exec(query, param...).Error
}

func (s *TransactionHistoryStore) TransactionHistoryBy(field string, value interface{}) ([]entity.TransactionHistory, error) {
	result := make([]entity.TransactionHistory, 0)
	err := s.DB.Where(map[string]interface{}{field: value}).Find(&result).Error
	return result, err
}
