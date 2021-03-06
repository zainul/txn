package store

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zainul/txn/internal/constant"
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

func (s *TransactionLogStore) Transfer(e []entity.TransactionLog) error {
	tx := s.DB.Begin()

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	// Insert to log
	for _, val := range e {
		txTime := time.Now()
		val.Date = txTime
		val.LastUpdate = txTime

		errInsert := tx.Create(&val).Error
		if errInsert != nil {
			tx.Rollback()
			return errInsert
		}
	}

	// Update user_account, log
	for _, val := range e {
		var user struct {
			EndingBalance float64 `gorm:"column:end_balance"`
		}

		var qUpdateUserAccount string
		var qUpdateTransactionLog string

		if val.DrCr == constant.DR {
			qUpdateTransactionLog = `update transaction_log 
				set 
					end_balance = end_balance - $1,
					transaction_status = $2 
				where tx_id = $3 and account_number = $4
				RETURNING end_balance`
			qUpdateUserAccount = `update user_account set balance = balance - $1 where account_number = $2`
		} else if val.DrCr == constant.CR {
			qUpdateTransactionLog = `update transaction_log 
				set 
					end_balance = end_balance + $1,
					transaction_status = $2 
				where tx_id = $3 and account_number = $4
				RETURNING end_balance`
			qUpdateUserAccount = `update user_account set balance = balance + $1 where account_number = $2`
		}

		errLog := tx.Raw(qUpdateTransactionLog, val.Amount, constant.TransactionStatusSuccess, val.TxID, val.AccountNumber).Scan(&user).Error

		if user.EndingBalance < 0 && val.AccountNumber != constant.UserSystemID {
			tx.Rollback()
			return errors.New("Ending balance must be greather than 0")
		}

		if errLog != nil {
			tx.Rollback()
			return errLog
		}

		// Insert transaction history
		th := s.makeHistory(val, user.EndingBalance)
		errTxHist := tx.Create(&th).Error

		if errTxHist != nil {
			tx.Rollback()
			return errTxHist
		}

		errAcc := tx.Exec(qUpdateUserAccount, val.Amount, val.AccountNumber).Error

		if errAcc != nil {
			tx.Rollback()
			return errAcc
		}
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *TransactionLogStore) makeHistory(e entity.TransactionLog, endBalance float64) entity.TransactionHistory {
	txTime := time.Now()
	th := entity.TransactionHistory{
		AccountNumber:     e.AccountNumber,
		OpponentAccount:   e.OpponentAccountNumber,
		Amount:            e.Amount,
		DrCr:              e.DrCr,
		EndBalance:        endBalance,
		Information:       e.Information,
		PreviousBalance:   e.PreviousBalance,
		TransactionCode:   e.TransactionCode,
		TransactionStatus: e.TransactionStatus,
		TransactionDate:   txTime,
		TxID:              e.TxID,
		UserComment:       e.Information,
	}

	return th
}
