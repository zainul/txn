package transaction

import "github.com/zainul/txn/internal/repository"

// Usecase ...
type Usecase struct {
	txHistoryStore   repository.TransactionHistory
	txLogStore       repository.TransactionLog
	userAccountStore repository.UserAccount
}

// NewTransaction ...
func NewTransaction(
	txHistoryStore repository.TransactionHistory,
	txLogStore repository.TransactionLog,
	userAccountStore repository.UserAccount,
) Usecase {
	return Usecase{
		txHistoryStore,
		txLogStore,
		userAccountStore,
	}
}
