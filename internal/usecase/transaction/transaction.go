package transaction

import (
	"errors"
	"fmt"

	"github.com/zainul/txn/internal/pkg/error/usecaseerror"
	"github.com/zainul/txn/internal/pkg/randomdigit"

	"github.com/zainul/txn/internal/contract"

	"github.com/zainul/txn/internal/constant"
	"github.com/zainul/txn/internal/entity"
	"github.com/zainul/txn/internal/pkg/error/deliveryerror"
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

// InternalTransfer is internal transfer method
func (t *Usecase) InternalTransfer(txnParam contract.TransactionRequest) (contract.TransactionResponse, *deliveryerror.Error) {
	txID, err := t.Transfer(txnParam, constant.TransferMemberToMemberCode)

	return contract.TransactionResponse{
		TransactionID: txID,
	}, err
}

// Transfer is transfer as general
func (t *Usecase) Transfer(transfer contract.TransactionRequest, transactionCode string) (string, *deliveryerror.Error) {
	fromAcc, toAcc, errFrom, errTo := t.ValidateUserAccount(transfer.FromAccountNumber, transfer.ToAccountNumber)

	if errFrom != nil {
		return "", deliveryerror.GetError(usecaseerror.InvalidFromAccountNumber, errFrom)
	}

	if errTo != nil {
		return "", deliveryerror.GetError(usecaseerror.InvalidToAccountNumber, errTo)
	}

	if !t.ValidBalancePreTransaction(fromAcc, transfer.Amount) {
		return "", deliveryerror.GetError(usecaseerror.InsufficientBalance, errors.New(usecaseerror.InsufficientBalance))
	}

	// Generate trx ID
	txID, err := randomdigit.GenerateRandomString(49)

	if err != nil {
		return "", deliveryerror.GetError(usecaseerror.NotFoundCode, err)
	}

	txLog := makeRule(fromAcc, toAcc, transactionCode, transfer.Amount, txID)

	err = t.TxLogRepo.Transfer(txLog)

	if err != nil {
		msgErr := deliveryerror.GetError(usecaseerror.FailedToTransfer, err)
		msgErr.ErrorMsg = fmt.Sprintf(msgErr.ErrorMsg, transfer.FromAccountNumber, transfer.ToAccountNumber)
		return "", msgErr
	}

	return txID, nil
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
