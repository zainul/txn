package transaction

import (
	"errors"

	"github.com/zainul/txn/internal/constant"
	"github.com/zainul/txn/internal/entity"
	"github.com/zainul/txn/internal/repository"
)

type userAccountChannel struct {
	AccountID int64
	User      *entity.UserAccount
	Err       error
}

// Usecase ...
type Usecase struct {
	TxHistoryRepo   repository.TransactionHistory
	TxLogRepo       repository.TransactionLog
	UserAccountRepo repository.UserAccount
}

// ValidateUserAccount is account available or not
func (t *Usecase) ValidateUserAccount(from int64, to int64) (fromAcc entity.UserAccount, toAcc entity.UserAccount, errFrom error, errTo error) {
	users := make(chan userAccountChannel)
	userAccs := []int64{from, to}

	for _, val := range userAccs {
		go func(userID int64) {
			user, err := t.UserAccountRepo.UserAccountBy("account_number", userID)

			users <- userAccountChannel{
				AccountID: userID,
				User:      &user[0],
				Err:       err,
			}
		}(val)
	}

	for range userAccs {
		select {
		case u := <-users:
			user := u.User

			if u.AccountID == from && user.AccountNumber != 0 {
				fromAcc = *user
				continue
			}
			if u.AccountID == from && user.AccountNumber == 0 {
				errFrom = errors.New("Invalid from account number")
				continue
			}
			if u.AccountID == from && u.Err != nil {
				errFrom = u.Err
				continue
			}
			if u.AccountID == to && user.AccountNumber != 0 {
				toAcc = *user
				continue
			}
			if u.AccountID == to && user.AccountNumber == 0 {
				errTo = errors.New("Invalid to account number")
				continue
			}
			if u.AccountID == to && u.Err != nil {
				errTo = u.Err
				continue
			}
		}
	}

	return
}

// ValidBalancePreTransaction use for pre check before the transaction happen
func (t *Usecase) ValidBalancePreTransaction(fromAcc entity.UserAccount, amount float64) bool {

	if fromAcc.AccountNumber == constant.UserSystemID {
		return true
	}

	return (fromAcc.Balance - amount) >= 0
}
