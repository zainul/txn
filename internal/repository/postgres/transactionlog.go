package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/zainul/txn/internal/entity"
)

type TransactionLogStore struct {
	DB *gorm.DB
}

// NewTransactionLogStore ...
func NewTransactionLogStore(conn *gorm.DB) *TransactionLogStore {
	return &TransactionLogStore{
		DB: conn,
	}
}

func (s *TransactionLogStore) Insert(e entity.TransactionLog) error {
	return s.DB.Create(&e).Error
}

func (s *TransactionLogStore) Update(e entity.TransactionLog) error {
	return s.DB.Table(e.TableName()).Where("tx_id = ?", e.TxID).Update(&e).Error
}

func (s *TransactionLogStore) UpdateBy(e entity.TransactionLog, param map[string]interface{}) error {
	return s.DB.Table(e.TableName()).Where(param).Update(&e).Error
}

func (s *TransactionLogStore) Delete(id interface{}) error {
	return nil
}

func (s *TransactionLogStore) RawQuery(sql string, result interface{}) error {
	return s.DB.Raw(sql).Find(result).Error
}

func (s *TransactionLogStore) RawExec(query string, param ...interface{}) error {
	return s.DB.Exec(query, param...).Error
}

func (s *TransactionLogStore) TransactionLogBy(field string, value interface{}) ([]entity.TransactionLog, error) {
	result := make([]entity.TransactionLog, 0)
	err := s.DB.Where(map[string]interface{}{field: value}).Find(&result).Error
	return result, err
}
