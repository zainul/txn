package usecase

import (
	"github.com/zainul/txn/internal/contract"
	"github.com/zainul/txn/internal/pkg/error/deliveryerror"
	"github.com/zainul/txn/internal/repository"
	"github.com/zainul/txn/internal/usecase/transaction"
	"github.com/zainul/txn/internal/usecase/user"
)

// Transaction ...
type Transaction interface {
	InternalTransfer(txnParam contract.TransactionRequest) (contract.TransactionResponse, *deliveryerror.Error)
	GetBalance(accountNumber int64) (contract.GetBalanceResponse, *deliveryerror.Error)
}

// User ...
type User interface {
	SeedUser(n int) error
	// GetBalance(accountNumber int64) (contract.GetBalanceResponse, error)
}

// NewUser ...
func NewUser(
	userAccountStore repository.UserAccount,
) *user.Usecase {
	return &user.Usecase{
		UserAccountRepo: userAccountStore,
	}
}

// NewTransaction ...
func NewTransaction(
	txHistoryStore repository.TransactionHistory,
	txLogStore repository.TransactionLog,
	userAccountStore repository.UserAccount,
) *transaction.Usecase {
	return &transaction.Usecase{
		TxHistoryRepo:   txHistoryStore,
		TxLogRepo:       txLogStore,
		UserAccountRepo: userAccountStore,
	}
}
