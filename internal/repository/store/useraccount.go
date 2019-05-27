package store

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zainul/txn/internal/entity"
)

type UserAccountStore struct {
	DB *gorm.DB
}

// NewUserAccountStore ...
func NewUserAccountStore(conn *gorm.DB) *UserAccountStore {
	return &UserAccountStore{
		DB: conn,
	}
}

func (s *UserAccountStore) Insert(e entity.UserAccount) error {
	now := time.Now()
	e.Modified = now
	return s.DB.Create(&e).Error
}

func (s *UserAccountStore) Update(e entity.UserAccount) error {
	return s.DB.Table(e.TableName()).Where("account_number = ?", e.AccountNumber).Update(&e).Error
}

func (s *UserAccountStore) UpdateBy(e entity.UserAccount, param map[string]interface{}) error {
	return s.DB.Table(e.TableName()).Where(param).Update(&e).Error
}

func (s *UserAccountStore) Delete(id interface{}) error {
	return nil
}

func (s *UserAccountStore) RawQuery(sql string, result interface{}) error {
	return s.DB.Raw(sql).Find(result).Error
}

func (s *UserAccountStore) RawExec(query string, param ...interface{}) error {
	return s.DB.Exec(query, param...).Error
}

func (s *UserAccountStore) UserAccountBy(field string, value interface{}) ([]entity.UserAccount, error) {
	result := make([]entity.UserAccount, 0)
	err := s.DB.Where(map[string]interface{}{field: value}).Find(&result).Error
	return result, err
}
